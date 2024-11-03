package database

import (
	"context"
 	"app_food/pkg/log"
 	"app_food/pkg/config"
	"time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"

)

type MongoInstance struct {
    Client *mongo.Client
    DB     *mongo.Database
}

var MI MongoInstance

func Connect() {
	connectionURI, _ := config.GetConfig("MONGO_CONNECTION_URI")
	database, _ := config.GetConfig("MONGO_DATABASE")

	if connectionURI == "" || database == ""  {
		log.Logger.Error("Missing MongoDB configuration in environment variables")
	}

    clientOptions := options.Client().ApplyURI(connectionURI)
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Logger.Error(err)
    }

    err = client.Ping(ctx, readpref.Primary())
    if err != nil {
        log.Logger.Error(err)
    }

	log.Logger.Info("Connect to MongoDB successfully!")

    MI = MongoInstance{
        Client: client,
        DB:     client.Database(database),
    }


}

func GetCollection(collectionName string) *mongo.Collection {
	if MI.DB == nil {
		log.Logger.Error("No MongoDB database is connected")
		return nil
	}
	return MI.DB.Collection(collectionName)
}

func init() {
	Connect()
}