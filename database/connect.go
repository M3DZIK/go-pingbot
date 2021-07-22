package database

import (
	"context"
	"os"
	"time"

	"gitlab.com/gaming0skar123/go/pingbot/common"
	"gitlab.com/gaming0skar123/go/pingbot/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client mongo.Client
var Coll *mongo.Collection

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Mongo_URI))
	if common.CheckErr(err, "connect to db") {
		os.Exit(1)
	}

	err = Client.Ping(ctx, readpref.Primary())
	if common.CheckErr(err, "ping db") {
		os.Exit(1)
	}

	Coll = Client.Database(config.Mongo_DB).Collection(config.Mongo_Collection)
}
