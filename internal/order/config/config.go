package config

import (
	"log"

	"github.com/google/wire"
	"github.com/pot-code/go-cqrs-demo/pkg/mq"
	"github.com/pot-code/gobit/pkg/config"
	"github.com/pot-code/gobit/pkg/db"
	"github.com/pot-code/gobit/pkg/logging"
	"github.com/pot-code/gobit/pkg/util"
	"github.com/pot-code/gobit/pkg/validate"
)

type AppConfig struct {
	Port          int32                   `validate:"required,min=1" mapstructure:"port" yaml:"port"`
	Base          *config.BaseConfig      `validate:"required" mapstructure:"base" yaml:"base"`
	Logging       *logging.LoggingConfig  `validate:"required" mapstructure:"logging" yaml:"logging"`
	Database      *db.DatabaseConfig      `validate:"required" mapstructure:"database" yaml:"database"`
	Security      *SecurityConfig         `validate:"required" mapstructure:"security" yaml:"security"`
	Kafka         *mq.KafkaProducerConfig `validate:"required" mapstructure:"kafka" yaml:"kafka"`
	EventTopics   *EventTopicConfig       `mapstructure:"event_topics" yaml:"event_topics"`
	EventSourcing *db.CacheConfig         `validate:"required" mapstructure:"event_sourcing" yaml:"event_sourcing"`
}

type EventTopicConfig struct {
	OrderCreated   string `validate:"required" mapstructure:"order_created" yaml:"order_created"`
	OrderConfirmed string `validate:"required" mapstructure:"order_confirmed" yaml:"order_confirmed"`
	OrderCanceled  string `validate:"required" mapstructure:"order_canceled" yaml:"order_canceled"`
}

type SecurityConfig struct {
	CORS []string `mapstructure:"cors" yaml:"cors"`
}

func NewAppConfig() *AppConfig {
	cfg := new(AppConfig)
	cm := config.NewConfigManager(config.WithConfigName("order"))
	util.HandleFatalError("failed to load config", cm.LoadConfig(cfg))

	v := validate.NewValidator()
	if err := v.Struct(cfg); err != nil {
		log.Fatalf("failed to validate config: \n%v", err)
	}
	return cfg
}

var ConfigSet = wire.NewSet(
	wire.FieldsOf(new(*AppConfig), "Base", "Logging", "Database", "Security", "Kafka", "EventTopics", "EventSourcing"),
	NewAppConfig,
)
