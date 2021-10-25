package service

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// This is to make sure all the orders are constructed and nothing is left out
func TestOrderWeightsConstructor(t *testing.T) {
	orders := createOrders()
	orderCollection := constructOrdersBasedOnWeightLimit(orders, 500)
	totalOrderIds := 0
	totalWeight := 0.0
	for _, order := range orderCollection {
		totalOrderIds += len(order.OrderIds)
		totalWeight += order.TotalWeight
	}
	assert.Equal(t, 1000, totalOrderIds)
	assert.Equal(t, 1000*24.45, math.Round(totalWeight*100)/100)
}

func createOrders() []Order {
	orders := make([]Order, 0)
	for i := 1; i <= 1000; i++ {
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
