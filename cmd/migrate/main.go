package main

import (
	"context"
	"time"

	"github.com/Shopify/sarama"
	"github.com/pot-code/go-cqrs-demo/migrate"
	"github.com/pot-code/gobit/pkg/util"
)

func main() {
	lm := util.NewLifecycleManager()
	defer lm.Exit(30 * time.Second)

	m := migrate.InitMigrate(lm)
	ctx := context.Background()
	util.HandleFatalError("failed creating schema resources", m.Client.Schema.Create(ctx))

	m.KafkaAdmin.CreateTopic("order", &sarama.TopicDetail{
		NumPartitions:     4,
		ReplicationFactor: 1,
	}, false)
}
