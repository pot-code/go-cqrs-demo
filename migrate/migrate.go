package migrate

import (
	"github.com/Shopify/sarama"
	"github.com/pot-code/go-cqrs-demo/ent"
)

type Migration struct {
	Client     *ent.Client
	KafkaAdmin sarama.ClusterAdmin
}

func NewMigration(client *ent.Client, admin sarama.ClusterAdmin) *Migration {
	return &Migration{Client: client, KafkaAdmin: admin}
}
