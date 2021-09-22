//go:build wireinject
// +build wireinject

package web

import (
	"github.com/google/wire"
	"github.com/pot-code/go-cqrs-demo/event"
	"github.com/pot-code/go-cqrs-demo/internal/order/command"
	"github.com/pot-code/go-cqrs-demo/internal/order/config"
	"github.com/pot-code/go-cqrs-demo/internal/order/domain"
	"github.com/pot-code/go-cqrs-demo/internal/order/query"
	repo "github.com/pot-code/go-cqrs-demo/internal/order/repository"
	pdb "github.com/pot-code/go-cqrs-demo/pkg/db"
	"github.com/pot-code/go-cqrs-demo/pkg/mq"
	"github.com/pot-code/go-cqrs-demo/pkg/validate"
	"github.com/pot-code/gobit/pkg/db"
	"github.com/pot-code/gobit/pkg/logging"
	"github.com/pot-code/gobit/pkg/util"
	"github.com/pot-code/gobit/pkg/uuid"
)

func InitHttpServer(em *util.LifecycleManager) *HttpServer {
	wire.Build(
		NewHttpServer, HttpSet, config.ConfigSet,
		command.CommandSet, query.QuerySet, mq.NewKafkaAsyncProducer, mq.NewKafkaPublisher,
		db.NewSqlxProvider, logging.NewZapLoggerProvider, pdb.NewEntClient,
		uuid.NewGoUUIDGenerator, validate.NewValidator,
		domain.NewOrderService,
		db.NewRedisCacheProvider,
		repo.NewOrderRedisRepository,
		wire.Bind(new(domain.OrderRepository), new(*repo.OrderRedisRepository)),
		// order.NewOrderPostgresRepository,
		// wire.Bind(new(order.OrderRepository), new(*order.OrderPostgresRepository)),
		wire.Bind(new(uuid.UUID), new(*uuid.GoUUIDGenerator)),
		wire.Bind(new(event.Publisher), new(*mq.KafkaPublisher)),
	)
	return &HttpServer{}
}
