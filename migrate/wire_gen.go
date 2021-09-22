// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package migrate

import (
	"github.com/pot-code/go-cqrs-demo/migrate/config"
	db2 "github.com/pot-code/go-cqrs-demo/pkg/db"
	"github.com/pot-code/go-cqrs-demo/pkg/mq"
	"github.com/pot-code/gobit/pkg/db"
	"github.com/pot-code/gobit/pkg/util"
)

// Injectors from wire.go:

func InitMigrate(lm *util.LifecycleManager) *Migration {
	appConfig := config.NewAppConfig()
	baseConfig := appConfig.Base
	databaseConfig := appConfig.Database
	sqlxDB := db.NewSqlxProvider(databaseConfig, lm)
	client := db2.NewEntClient(baseConfig, sqlxDB, lm)
	kafkaAdminConfig := appConfig.KafkaAdmin
	clusterAdmin := mq.NewKafkaAdmin(kafkaAdminConfig, lm)
	migration := NewMigration(client, clusterAdmin)
	return migration
}
