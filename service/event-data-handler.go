package service

import (
	"encoding/json"
	"khidr/service-b/persistence"
	"time"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Id          int     `json:"courier_id"`
	Email       string  `json:"email"`
	PhoneNumber string  `json:"phone_number"`
	Weight      float64 `json:"weight"`
}

func (s *Service) RunDataHandler() {
	messages := make(chan []byte)
	go s.KafkaConfig.RunConsumer(messages)
	go func(messages chan []byte) {
		for {
			select {
			case val, ok := <-messages:
				if ok {
					orders, err := parseEventMessages(val)
					if err != nil {
						log.Error(err)
					}
					s.storeOrdersToCollectionByCountry(orders)
				}
			}

		}
	}(messages)
}

func parseEventMessages(messages []byte) (map[string][]Order, error) {
	var orders map[string][]Order
	err := json.Unmarshal(messages, &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *Service) storeOrdersToCollectionByCountry(orders map[string][]Order) {
	if cameroon, ok := orders["Cameroon"]; ok {
		cameroonData := parseCameroonCollection(cameroon)
		s.MongoConfig.BulkInsertCameroonOrders(cameroonData)
	}
}

func parseCameroonCollection(cameroonOrders []Order) []persistence.CameroonOrder {
	cameroonData := make([]persistence.CameroonOrder, 0)
	for _, order := range cameroonOrders {
		cameroonOrder := persistence.CameroonOrder{
			Id:        primitive.NewObjectID(),
			CreatedAt: time.Now(),
			OrderId:   order.Id,
			Weight:    order.Weight,
		}
		cameroonData = append(cameroonData, cameroonOrder)
	}
	return cameroonData
}
