package persistence

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoConfig struct {
	ctx      context.Context
	uri      string
	database string
	client   *mongo.Client
}

func InitMongo(ctx context.Context, uri, database string) *MongoConfig {
	return &MongoConfig{
		ctx:      ctx,
		uri:      uri,
		database: database,
	}
}

func (m *MongoConfig) Connect() error {
	credential := options.Credential{
		Username: "admin",
		Password: "admin",
	}
	clientOpts := options.Client().ApplyURI(m.uri).SetAuth(credential)
	client, err := mongo.Connect(m.ctx, clientOpts)
	if err != nil {
		return err
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return err
	}
	m.client = client
	return nil
}

func (m *MongoConfig) Disconnect() error {
	if err := m.client.Disconnect(context.TODO()); err != nil {
		return err
	}
	return nil
}

func (m *MongoConfig) CreateIndexes() error {
	collections := []string{"Cameroon", "Ethiopia", "Morocco", "Mozambique", "Uganda"}
	for _, collection := range collections {
		if err := m.createIndexesOnCreatedAtField(collection); err != nil {
			return err
		}
	}
	return nil
}

func (m *MongoConfig) createIndexesOnCreatedAtField(collection string) error {
	coll := m.client.Database(m.database).Collection(collection)
	_, err := coll.Indexes().CreateOne(
		m.ctx,
		mongo.IndexModel{
			Keys: bson.D{
				{Key: "created_at", Value: 1},
			},
		},
	)
	return err
}
