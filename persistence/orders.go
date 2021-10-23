package persistence

import (
	"time"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Order struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	OrderIds    []int              `json:"order_ids" bson:"order_ids"`
	TotalWeight float64            `json:"total_weight" bson:"total_weight"`
}

func (m *MongoConfig) BulkInsertOrders(collection string, orders []Order) {
	coll := m.client.Database(m.database).Collection(collection)
	var operations []mongo.WriteModel
	for i := range orders {
		order := orders[i]
		updateFilter := bson.M{"_id": order.Id}
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
