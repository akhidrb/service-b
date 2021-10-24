package events

import (
	"context"

	"github.com/labstack/gommon/log"
	"github.com/segmentio/kafka-go"
)

type KafkaConfig struct {
	ctx             context.Context
	bootstrapServer string
	consumerGroup   string
	topic           string
}

func InitKafkaConfig(ctx context.Context, bootstrapServer, consumerGroupId, topic string) *KafkaConfig {
	return &KafkaConfig{
		ctx:             ctx,
		bootstrapServer: bootstrapServer,
		consumerGroup:   consumerGroupId,
		topic:           topic,
	}
}

func (k *KafkaConfig) RunConsumer(messages chan []byte) error {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{k.bootstrapServer},
		GroupID:  k.consumerGroup,
		Topic:    k.topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	defer r.Close()
	for {
		m, err := r.ReadMessage(k.ctx)
		if err != nil {
			log.Error(err)
		} else {
			messages <- m.Value
		}
	}
}
