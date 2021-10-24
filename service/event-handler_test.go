package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeightMapping(t *testing.T) {
	orders := createOrders()
	orderCollection := parseToCollectionModel(orders)
	totalOrderIds := 0
	for _, order := range orderCollection {
		totalOrderIds += len(order.OrderIds)
	}
	assert.Equal(t, 200, totalOrderIds)
}

func createOrders() []Order {
	orders := make([]Order, 0)
	for i := 1; i <= 200; i++ {
		order := Order{
			Id:          i,
			Email:       "email@email.com",
			PhoneNumber: "237 209993809",
			Weight:      24.45,
		}
		orders = append(orders, order)
	}
	return orders
}
