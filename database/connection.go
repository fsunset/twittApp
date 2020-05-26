package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConnection initializes connection
var MongoConnection = databaseConnection()
var clientOptions = options.Client().ApplyURI("mongodb+srv://twittAppClusterUser:twittAppClusterPass20@twittappcluster-6ibps.mongodb.net/test?retryWrites=true&w=majority")

func databaseConnection() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("Error on DB Init Connection : " + err.Error())
		return client
	}

	log.Println("MongoDB successfully connected")
	return client
}

// CheckConnection checks mongoDB connection
func CheckConnection() bool {
	err := MongoConnection.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal("Error on DB CheckConnection : " + err.Error())
		return false
	}

	return true
}
