package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/pot-code/go-cqrs-demo/internal/writer"
	"github.com/pot-code/gobit/pkg/util"
	"go.uber.org/zap"
)

func main() {
	lm := util.NewLifecycleManager()
	defer lm.WaitExitSignal(30 * time.Second)

	w := writer.InitWriter(lm)
	logger := w.Logger
	ctx, cancel := context.WithCancel(context.Background())
	wg := new(sync.WaitGroup)
	topics := w.Config.KafkaConsumer.Topics

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			util.HandlePanicError("kafka failed to consume",
				w.KafkaConsumerGroup.Consume(ctx, topics, w.OrderConsumer))

			if ctx.Err() != nil {
				return
			}

			w.OrderConsumer.Ready = make(chan bool)
		}
	}()

	<-w.OrderConsumer.Ready

	logger.Info("writer service is up", zap.Strings("topics", topics),
		zap.String("consumer_group", w.Config.KafkaConsumer.ConsumerGroup))
	lm.OnExit(func(_ context.Context) {
		cancel()
		wg.Wait()
		w.KafkaConsumerGroup.Close()
		log.Println("[sarama.ConsumerGroup] close consumer group")
	})
}
