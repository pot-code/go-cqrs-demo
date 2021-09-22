package web

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Shopify/sarama"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	echo_middleware "github.com/labstack/echo/v4/middleware"
	"github.com/pot-code/go-cqrs-demo/internal/order/config"
	"github.com/pot-code/go-cqrs-demo/internal/order/port"
	gobit "github.com/pot-code/gobit/pkg"
	"github.com/pot-code/gobit/pkg/api"
	"github.com/pot-code/gobit/pkg/logging"
	"github.com/pot-code/gobit/pkg/middleware"
	"github.com/pot-code/gobit/pkg/util"
	"go.uber.org/zap"
)

type HandlerCollection struct {
	*port.OrderHandler
}

type HttpServer struct {
	Config        *config.AppConfig
	Logger        *zap.Logger
	Server        *echo.Echo
	Handlers      *HandlerCollection
	KafkaProducer sarama.AsyncProducer
}

func NewHttpServer(
	cfg *config.AppConfig,
	logger *zap.Logger,
	server *echo.Echo,
	handlers *HandlerCollection,
	p sarama.AsyncProducer,
) *HttpServer {
	return &HttpServer{Config: cfg, Logger: logger, Server: server, Handlers: handlers, KafkaProducer: p}
}

func NewEchoServer(cfg *config.AppConfig, logger *zap.Logger, lm *util.LifecycleManager) *echo.Echo {
	app := echo.New()

	if cfg.Base.Env == gobit.EnvProduction {
		app.HideBanner = true
	}

	app.GET("/healthz", func(c echo.Context) error {
		return lm.Probe(5 * time.Second)
	})

	app.Use(
		middleware.ErrorHandling(middleware.ErrorHandlingOption{Handler: func(c echo.Context, err error) {
			msg := api.ErrInternalError.Error()
			logger.Error(msg, zap.Object("error", logging.NewZapStacktraceError(err, 3)))
			c.JSON(http.StatusInternalServerError,
				api.NewRESTStandardError(msg),
			)
		}}),
		echo_middleware.Gzip(),
		echo_middleware.CORS(),
		echo_middleware.Secure(),
		middleware.ParseAcceptLanguage(middleware.ParseAcceptLanguageOption{}),
	)

	lm.OnExit(func(ctx context.Context) {
		log.Println("[echo.Echo] shutdown server")
		app.Close()
	})

	return app
}

var HttpSet = wire.NewSet(
	wire.Struct(new(HandlerCollection), "*"),
	NewEchoServer,
	port.NewOrderHandler,
)
