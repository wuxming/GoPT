package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Database  *mongo.Database
func Coll() *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client,err := mongo.Connect(context.TODO(),clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(),nil)
	if err != nil {
		log.Fatal(err)
	}
	Database = client.Database("ittest")
	return Database
}
