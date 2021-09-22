package mq

import (
	"context"
	"log"

	"github.com/Shopify/sarama"
	"github.com/pot-code/gobit/pkg/util"
	"go.uber.org/zap/zapcore"
)

type KafkaProducerConfig struct {
	Brokers []string `mapstructure:"brokers" yaml:"brokers"`
}

func NewKafkaAsyncProducer(kc *KafkaProducerConfig, lm *util.LifecycleManager) sarama.AsyncProducer {
	if kc == nil {
		panic("KafkaProducerConfig is nil")
	}

	config := sarama.NewConfig()
	config.Version = sarama.V2_0_0_0
	p, err := sarama.NewAsyncProducer(kc.Brokers, config)
	util.HandlePanicError("failed to init kafka producer", err)

	lm.OnExit(func(ctx context.Context) {
		log.Println("[sarama.AsyncProducer] close producer")
		p.Close()
	})

	return p
}

type KafkaConsumerConfig struct {
	Brokers       []string `mapstructure:"brokers" yaml:"brokers"`
	Topics        []string `validate:"required" mapstructure:"topics" yaml:"topics"`
	ConsumerGroup string   `validate:"required" mapstructure:"consumer_group" yaml:"consumer_group"`
}

func NewKafkaConsumerGroup(kc *KafkaConsumerConfig, lm *util.LifecycleManager) sarama.ConsumerGroup {
	if kc == nil {
		panic("KafkaConsumerConfig is nil")
	}

	config := sarama.NewConfig()
	config.Version = sarama.V2_0_0_0
	c, err := sarama.NewConsumerGroup(kc.Brokers, kc.ConsumerGroup, config)
	util.HandlePanicError("failed to init kafka consumer", err)
	return c
}

type KafkaAdminConfig struct {
	Brokers []string `mapstructure:"brokers" yaml:"brokers"`
}

func NewKafkaAdmin(kc *KafkaAdminConfig, lm *util.LifecycleManager) sarama.ClusterAdmin {
	if kc == nil {
		panic("KafkaAdminConfig is nil")
	}

	config := sarama.NewConfig()
	config.Version = sarama.V2_0_0_0
	admin, err := sarama.NewClusterAdmin(kc.Brokers, config)
	util.HandlePanicError("failed to init kafka cluster admin", err)

	lm.OnExit(func(ctx context.Context) {
		log.Println("[sarama.ClusterAdmin] close admin")
		admin.Close()
	})
	return admin
}

type KafkaZapProducerMessage struct {
	msg *sarama.ProducerMessage
}

func NewKafkaZapProducerMessage(msg *sarama.ProducerMessage) *KafkaZapProducerMessage {
	return &KafkaZapProducerMessage{msg}
}

func (e *KafkaZapProducerMessage) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	msg := e.msg

	et := FromHeaders(msg.Headers).GetType()
	enc.AddString("topic", msg.Topic)
	enc.AddInt32("partition", msg.Partition)
	enc.AddInt64("offset", msg.Offset)
	enc.AddString("msg_type", et)
	return nil
}

type KafkaZapConsumerMessage struct {
	msg *sarama.ConsumerMessage
}

func NewKafkaZapConsumerMessage(msg *sarama.ConsumerMessage) *KafkaZapConsumerMessage {
	return &KafkaZapConsumerMessage{msg}
}

func (kcm *KafkaZapConsumerMessage) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	msg := kcm.msg

	et := FromPointerHeaders(msg.Headers).GetType()
	enc.AddString("topic", msg.Topic)
	enc.AddInt32("partition", msg.Partition)
	enc.AddInt64("offset", msg.Offset)
	enc.AddString("msg_type", et)
	return nil
}
