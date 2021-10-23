package events

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/labstack/gommon/log"
)

type KafkaConfig struct {
	BootstrapServer string
	ConsumerGroupId string
	Topic           string
}

func InitKafkaConfig(bootstrapServer, consumerGroupId, topic string) *KafkaConfig {
	return &KafkaConfig{
		BootstrapServer: bootstrapServer,
		ConsumerGroupId: consumerGroupId,
		Topic:           topic,
	}
}

func (k *KafkaConfig) RunConsumer(messages chan []byte) error {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": k.BootstrapServer,
		"group.id":          k.ConsumerGroupId,
		"auto.offset.reset": "latest",
	})
	if err != nil {
		return err
	}
	defer c.Close()
	c.SubscribeTopics([]string{k.Topic}, nil)
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			messages <- msg.Value
		} else {
			log.Error(err)
		}
	}
}
