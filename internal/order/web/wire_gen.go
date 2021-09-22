// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package web

import (
	"github.com/pot-code/go-cqrs-demo/internal/order/command"
	"github.com/pot-code/go-cqrs-demo/internal/order/config"
	"github.com/pot-code/go-cqrs-demo/internal/order/domain"
	"github.com/pot-code/go-cqrs-demo/internal/order/port"
	"github.com/pot-code/go-cqrs-demo/internal/order/query"
	"github.com/pot-code/go-cqrs-demo/internal/order/repository"
	db2 "github.com/pot-code/go-cqrs-demo/pkg/db"
	"github.com/pot-code/go-cqrs-demo/pkg/mq"
	"github.com/pot-code/go-cqrs-demo/pkg/validate"
	"github.com/pot-code/gobit/pkg/db"
	"github.com/pot-code/gobit/pkg/logging"
	"github.com/pot-code/gobit/pkg/util"
	"github.com/pot-code/gobit/pkg/uuid"
)

// Injectors from wire.go:

func InitHttpServer(em *util.LifecycleManager) *HttpServer {
	appConfig := config.NewAppConfig()
	loggingConfig := appConfig.Logging
	logger := logging.NewZapLoggerProvider(loggingConfig, em)
	echo := NewEchoServer(appConfig, logger, em)
	kafkaProducerConfig := appConfig.Kafka
	asyncProducer := mq.NewKafkaAsyncProducer(kafkaProducerConfig, em)
	kafkaPublisher := mq.NewKafkaPublisher(asyncProducer)
	cacheConfig := appConfig.EventSourcing
	client := db.NewRedisCacheProvider(cacheConfig, em)
	orderRedisRepository := repo.NewOrderRedisRepository(client)
	eventTopicConfig := appConfig.EventTopics
	orderService := domain.NewOrderService(kafkaPublisher, orderRedisRepository, eventTopicConfig)
	createOrderHandler := command.NewCreateOrderHandler(logger, orderService)
	deleteOrderHandler := command.NewDeleteOrderHandler(logger, kafkaPublisher)
	confirmOrderHandler := command.NewConfirmOrderHandler(logger, orderService)
	cancelOrderHandler := command.NewCancelOrderHandler(logger)
	orderCommands := &command.OrderCommands{
		CreateOrderHandler:  createOrderHandler,
		DeleteOrderHandler:  deleteOrderHandler,
		ConfirmOrderHandler: confirmOrderHandler,
		CancelOrderHandler:  cancelOrderHandler,
	}
	baseConfig := appConfig.Base
	databaseConfig := appConfig.Database
	sqlxDB := db.NewSqlxProvider(databaseConfig, em)
	entClient := db2.NewEntClient(baseConfig, sqlxDB, em)
	queryByIdHandler := query.NewQueryByIdHandler(entClient)
	orderQueries := &query.OrderQueries{
		QueryByIdHandler: queryByIdHandler,
	}
	goUUIDGenerator := uuid.NewGoUUIDGenerator()
	validatorV10 := validate.NewValidator()
	orderHandler := port.NewOrderHandler(orderCommands, orderQueries, goUUIDGenerator, validatorV10)
	handlerCollection := &HandlerCollection{
		OrderHandler: orderHandler,
	}
	httpServer := NewHttpServer(appConfig, logger, echo, handlerCollection, asyncProducer)
	return httpServer
}