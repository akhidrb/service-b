package service

import (
	"khidr/service-b/events"
	"khidr/service-b/persistence"
)

type Service struct {
	KafkaConfig *events.KafkaConfig
	MongoConfig *persistence.MongoConfig
}

func New(kafkaConfig *events.KafkaConfig, mongoConfig *persistence.MongoConfig) (*Service, error) {
	return &Service{
		KafkaConfig: kafkaConfig,
		MongoConfig: mongoConfig,
	}, nil
}
