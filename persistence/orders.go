package persistence

import (
	"time"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CameroonOrder struct {
	Id        primitive.ObjectID `json:"id"`
	CreatedAt time.Time          `json:"created_at"`
	OrderId   int                `json:"order_id"`
	Weight    float64            `json:"weight"`
}

func (m *MongoConfig) BulkInsertCameroonOrders(orders []CameroonOrder) {
	coll := m.client.Database(m.database).Collection(m.cameroonCollection)
	var operations []mongo.WriteModel
	for i := range orders {
		order := orders[i]
		updateFilter := bson.M{"order_id": order.OrderId}
		updateOperation := bson.M{
			"$set": order,
		}
		updateModel := mongo.NewUpdateOneModel()
		updateModel.SetFilter(updateFilter)
		updateModel.SetUpdate(updateOperation)
		updateModel.SetUpsert(true)
		operations = append(operations, updateModel)
	}
	bulkOptions := options.BulkWriteOptions{}
	bulkOptions.SetOrdered(false)
	_, err := coll.BulkWrite(m.ctx, operations, &bulkOptions)
	if err != nil {
		log.Error(err)
	}
}
