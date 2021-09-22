//+build wireinject

package writer

import (
	"github.com/Shopify/sarama"
	"github.com/google/wire"
	"github.com/pot-code/go-cqrs-demo/internal/writer/config"
	"github.com/pot-code/go-cqrs-demo/internal/writer/consumer"
	pdb "github.com/pot-code/go-cqrs-demo/pkg/db"
	"github.com/pot-code/go-cqrs-demo/pkg/mq"
	"github.com/pot-code/gobit/pkg/db"
	"github.com/pot-code/gobit/pkg/logging"
	"github.com/pot-code/gobit/pkg/util"
	"go.uber.org/zap"
)

type HandlerCollections struct {
	*consumer.OrderCreateHandler
	*consumer.OrderConfirmHandler
}

type Writer struct {
	KafkaConsumerGroup sarama.ConsumerGroup
	OrderConsumer      *mq.KafkaEventConsumer
	Config             *config.AppConfig
	Logger             *zap.Logger
}

func NewWriter(
	kc sarama.ConsumerGroup,
	cfg *config.AppConfig,
	c *mq.KafkaEventConsumer,
	hc *HandlerCollections,
	logger *zap.Logger,
) *Writer {
	return &Writer{KafkaConsumerGroup: kc, OrderConsumer: c, Logger: logger, Config: cfg}
}

var HandlerSet = wire.NewSet(
	wire.Struct(new(HandlerCollections), "*"),
	consumer.NewOrderCreateHandler,
	consumer.NewOrderConfirmHandler,
)

func InitWriter(lm *util.LifecycleManager) *Writer {
	wire.Build(
		NewWriter, mq.NewKafkaEventConsumer,
		config.ConfigSet, HandlerSet,
		mq.NewKafkaConsumerGroup,
		logging.NewZapLoggerProvider,
		pdb.NewEntClient,
		db.NewSqlxProvider,
	)
	return &Writer{}
}
