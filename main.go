package main

import (
	"khidr/service-b/api"
	"khidr/service-b/events"
	"khidr/service-b/persistence"
	"khidr/service-b/service"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

type Config struct {
	ListenAddress    string
	Timeout          int
	KafkaServer      string
	KafkaTopic       string
	ConsumerGroup    string
	MongoURI         string
	DatabaseName     string
	CargoWeightLimit float64
}

func main() {
	conf := Config{}
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Name = "Service B"
	app.Usage = "API for Retrieving Daily Manifest File"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "port",
			Usage:       "`IP:PORT` for the service to listen on",
			Value:       ":8081",
			Destination: &conf.ListenAddress,
		},
		&cli.IntFlag{
			Name:        "timeout",
			Usage:       "Requests `TIMEOUT` in seconds",
			Value:       30,
			Destination: &conf.Timeout,
		},
		&cli.StringFlag{
			Name:        "kafka-server",
			Usage:       "KafkaServer Host",
			Value:       "localhost",
			Destination: &conf.KafkaServer,
		},
		&cli.StringFlag{
			Name:        "kafka-topic",
			Usage:       "Kafka Topic to produce to",
			Value:       "test1",
			Destination: &conf.KafkaTopic,
		},
		&cli.StringFlag{
			Name:        "consumer-group",
			Usage:       "Kafka Consumer group",
			Value:       "cargoGroup",
			Destination: &conf.ConsumerGroup,
		},
		&cli.StringFlag{
			Name:        "mongo-uri",
			Usage:       "MongoDB URI",
			Value:       "mongodb://localhost:27017",
			Destination: &conf.MongoURI,
		},
		&cli.StringFlag{
			Name:        "db-name",
			Usage:       "Mongo Database Name",
			Value:       "cargos",
			Destination: &conf.DatabaseName,
		},
		&cli.Float64Flag{
			Name:        "cargo-weight-limit",
			Usage:       "Cargo Weight limit",
			Value:       500,
			Destination: &conf.CargoWeightLimit,
		},
	}

	app.Action = func(context *cli.Context) error {
		timeout := time.Duration(conf.Timeout) * time.Second
		kafkaConfig := events.InitKafkaConfig(context.Context, conf.KafkaServer, conf.ConsumerGroup, conf.KafkaTopic)
		mongo := persistence.InitMongo(context.Context, conf.MongoURI, conf.DatabaseName)
		err := mongo.Connect()
		if err != nil {
			return err
		}
		err = mongo.CreateIndexes()
		if err != nil {
			return err
		}
		defer func() {
			if err := mongo.Disconnect(); err != nil {
				log.Fatal(err)
			}
		}()
		serviceB, err := service.New(kafkaConfig, mongo, conf.CargoWeightLimit)
		if err != nil {
			return err
		}
		serviceB.RunDataHandler()
		server := api.New(conf.ListenAddress, timeout, serviceB)
		log.Println("Service B Running on Port " + conf.ListenAddress)
		return server.Start()
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
