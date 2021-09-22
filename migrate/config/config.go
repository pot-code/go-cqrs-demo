package config

import (
	"log"

	"github.com/google/wire"
	"github.com/pot-code/go-cqrs-demo/pkg/mq"
	"github.com/pot-code/gobit/pkg/config"
	"github.com/pot-code/gobit/pkg/db"
	"github.com/pot-code/gobit/pkg/util"
	"github.com/pot-code/gobit/pkg/validate"
)

type AppConfig struct {
	Base       *config.BaseConfig   `validate:"required" mapstructure:"base" yaml:"base"`
	Database   *db.DatabaseConfig   `validate:"required" mapstructure:"database" yaml:"database"`
	KafkaAdmin *mq.KafkaAdminConfig `validate:"required" mapstructure:"kafka_admin" yaml:"kafka_admin"`
}

func NewAppConfig() *AppConfig {
	cfg := new(AppConfig)
	cm := config.NewConfigManager(config.WithConfigName("migrate"))
	util.HandleFatalError("failed to load config", cm.LoadConfig(cfg))

	v := validate.NewValidator()
	if err := v.Struct(cfg); err != nil {
		log.Fatalf("failed to validate config: \n%v", err)
	}
	return cfg
}

var ConfigSet = wire.NewSet(
	wire.FieldsOf(new(*AppConfig), "Base", "Database", "KafkaAdmin"),
	NewAppConfig,
)
