package consumer

import (
	"context"
	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
	"github.com/pot-code/go-cqrs-demo/ent"
	"github.com/pot-code/go-cqrs-demo/event"
	"github.com/pot-code/go-cqrs-demo/pkg/mq"
	"go.uber.org/zap"
)

type OrderConfirmHandler struct {
	logger *zap.Logger
	ec     *ent.Client
}

var _ mq.EventHandler = &OrderConfirmHandler{}

func NewOrderConfirmHandler(
	c *mq.KafkaEventConsumer,
	logger *zap.Logger,
	ec *ent.Client,
) *OrderConfirmHandler {
	self := &OrderConfirmHandler{logger, ec}
	c.RegisterHandler(new(event.OrderConfirmedEvent), self)
	return self
}

func (h *OrderConfirmHandler) Handle(ctx context.Context, msg *sarama.ConsumerMessage) error {
	e := new(event.OrderConfirmedEvent)
	if err := json.Unmarshal(msg.Value, e); err != nil {
		return errors.Wrap(err, "failed to unmarshal OrderConfirmedEvent")
	}

	_, err := h.ec.Order.UpdateOneID(e.ID).SetStatus(e.Status).Save(ctx)
	return errors.Wrap(err, "failed to update order")
}
