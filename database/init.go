package database

import (
	"context"
	"time"

	"gitlab.com/gaming0skar123/go/pingbot/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client
var DB *mongo.Database

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Mongo_URI))
	if err != nil {
		panic("Failed Connecting to DB: " + err.Error())
	}

	err = Client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic("Failed Pinging DB: " + err.Error())
	}

	DB = Client.Database(config.Mongo_DB)
}
