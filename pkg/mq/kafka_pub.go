package mq

import (
	"context"
	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/pot-code/go-cqrs-demo/event"
)

type KafkaPublisher struct {
	p sarama.AsyncProducer
}

var _ event.Publisher = &KafkaPublisher{}

func NewKafkaPublisher(p sarama.AsyncProducer) *KafkaPublisher {
	return &KafkaPublisher{p}
}

func (kp *KafkaPublisher) Publish(ctx context.Context, topic, key string, e event.Event) error {
	headers := NewKafkaHeaders().SetType(e.Name())
	jm, _ := json.Marshal(e)
	kp.p.Input() <- &sarama.ProducerMessage{
		Topic:   topic,
		Key:     sarama.ByteEncoder([]byte(key)),
		Value:   sarama.ByteEncoder(jm),
		Headers: headers.ExportHeaders(),
	}
	return nil
}
