package dao

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MgoBaseDao struct {
}

var onceDB sync.Once

var dbClient *mongo.Client

func getDBClient() *mongo.Client {
	onceDB.Do(func() {

		clientOptions := options.Client().ApplyURI("mongodb+srv://Hao:Hao12345678@huangcluster.fmtgp.mongodb.net")
		var ctx = context.TODO()
		// Connect to MongoDB
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}
		// Check the connection
		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}
		dbClient = client

	})
	return dbClient
}

func getCollection(collectionName string) *mongo.Collection {

	client := getDBClient()

	collection := client.Database("HuangcoolyDB").Collection(collectionName)

	return collection
}
