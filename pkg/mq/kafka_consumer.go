package mq

import (
	"context"

	"github.com/Shopify/sarama"
	"github.com/pot-code/go-cqrs-demo/event"
	"github.com/pot-code/gobit/pkg/config"
	"go.uber.org/zap"
)

type EventHandler interface {
	Handle(ctx context.Context, msg *sarama.ConsumerMessage) error
}

type KafkaEventConsumer struct {
	handlers map[string]EventHandler
	logger   *zap.Logger
	cfg      *config.BaseConfig
	Ready    chan bool
}

func NewKafkaEventConsumer(logger *zap.Logger, cfg *config.BaseConfig) *KafkaEventConsumer {
	return &KafkaEventConsumer{
		logger:   logger,
		handlers: make(map[string]EventHandler),
		cfg:      cfg,
		Ready:    make(chan bool),
	}
}

func (c *KafkaEventConsumer) RegisterHandler(e event.Event, handler EventHandler) {
	c.handlers[e.Name()] = handler
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (c *KafkaEventConsumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(c.Ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (c *KafkaEventConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (c *KafkaEventConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	handlers := c.handlers
	logger := c.logger

	for message := range claim.Messages() {
		et := FromPointerHeaders(message.Headers).GetType()
		if handler, ok := handlers[et]; ok {
			logger.Debug("consume message", zap.Object("kafka", NewKafkaZapConsumerMessage(message)))
			if err := handler.Handle(context.Background(), message); err != nil {
				logger.Error("failed to handle message", zap.Error(err), zap.Object("kafka", NewKafkaZapConsumerMessage(message)))
			}
		}
		session.MarkMessage(message, "")
	}

	return nil
}
