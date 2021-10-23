package persistence

import (
	"context"

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
