package service

import (
	"encoding/json"
	"khidr/service-b/persistence"
	"time"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Id          int     `json:"order_id"`
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
	countries := []string{"Cameroon", "Ethiopia", "Morocco", "Mozambique", "Uganda"}
	for _, country := range countries {
		go func(country string) {
			if ordersList, ok := orders[country]; ok {
				parsedData := constructOrdersBasedOnWeights(ordersList)
				s.MongoConfig.BulkInsertOrders(country, parsedData)
			}
		}(country)
	}
}

func constructOrdersBasedOnWeights(orders []Order) []persistence.Order {
	orderData := make([]persistence.Order, 0)
	var orderModel *persistence.Order
	for i := range orders {
		order := orders[i]
		if orderModel == nil {
			orderModel = &persistence.Order{
				Id:          primitive.NewObjectID(),
				CreatedAt:   time.Now(),
				TotalWeight: 0,
			}
		}
		if orderModel.TotalWeight+order.Weight <= 500 {
			orderModel.TotalWeight += order.Weight
			orderModel.OrderIds = append(orderModel.OrderIds, order.Id)
		} else {
			orderData = append(orderData, *orderModel)
			orderModel = nil
			orderModel = &persistence.Order{
				Id:          primitive.NewObjectID(),
				CreatedAt:   time.Now(),
				OrderIds:    []int{order.Id},
				TotalWeight: order.Weight,
			}
		}
	}
	orderData = append(orderData, *orderModel)
	return orderData
}
