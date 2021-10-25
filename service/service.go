package service

import (
	"khidr/service-b/events"
	"khidr/service-b/persistence"
)

type Service struct {
	KafkaConfig      *events.KafkaConfig
	MongoConfig      *persistence.MongoConfig
	CargoWeightLimit float64
}

func New(kafkaConfig *events.KafkaConfig, mongoConfig *persistence.MongoConfig, cargoWeightLimit float64) (*Service, error) {
	return &Service{
		KafkaConfig:      kafkaConfig,
		MongoConfig:      mongoConfig,
		CargoWeightLimit: cargoWeightLimit,
	}, nil
}
