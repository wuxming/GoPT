package initialization

import (
	"context"
	"gin1/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func MongoInit() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	MongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println(err)
	}
	err = MongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
	}
	ittest := MongoClient.Database("ittest")
	service.GetCustomersCollections(ittest)
}
