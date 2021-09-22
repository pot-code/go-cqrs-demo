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
	Base          *config.BaseConfig      `mapstructure:"base" yaml:"base"`
	Logging       *logging.LoggingConfig  `validate:"required" mapstructure:"logging" yaml:"logging"`
	Database      *db.DatabaseConfig      `validate:"required" mapstructure:"database" yaml:"database"`
	KafkaConsumer *mq.KafkaConsumerConfig `validate:"required" mapstructure:"kafka_consumer" yaml:"kafka_consumer"`
	KafkaProducer *mq.KafkaProducerConfig `validate:"required" mapstructure:"kafka_producer" yaml:"kafka_producer"`
}

func NewAppConfig() *AppConfig {
	cfg := new(AppConfig)
	cm := config.NewConfigManager(config.WithConfigName("writer"))
	util.HandleFatalError("failed to load config", cm.LoadConfig(cfg))

	v := validate.NewValidator()
	if err := v.Struct(cfg); err != nil {
		log.Fatalf("failed to validate config: \n%s", err)
	}
	return cfg
}

var ConfigSet = wire.NewSet(
	wire.FieldsOf(new(*AppConfig), "Base", "Logging", "Database", "KafkaConsumer", "KafkaProducer"),
	NewAppConfig,
)
