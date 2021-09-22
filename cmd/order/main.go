package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pot-code/go-cqrs-demo/internal/order/web"
	"github.com/pot-code/go-cqrs-demo/pkg/mq"
	"github.com/pot-code/gobit/pkg/api"
	"github.com/pot-code/gobit/pkg/util"
	"go.uber.org/zap"
)

func main() {
	lm := util.NewLifecycleManager()
	defer lm.WaitExitSignal(30 * time.Second)

	hs := web.InitHttpServer(lm)
	ep := web.NewEndpoint(hs)

	api.ApplyEndpoint(hs.Server, ep)
	api.PrintRoutes(hs.Server, hs.Logger)

	go func() {
		logger := hs.Logger
		for err := range hs.KafkaProducer.Errors() {
			logger.Error(err.Error(), zap.Object("kafka", mq.NewKafkaZapProducerMessage(err.Msg)))
		}
	}()

	go func() {
		if err := hs.Server.Start(fmt.Sprintf(":%d", hs.Config.Port)); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Fatal(err)
			}
		}
	}()
}
