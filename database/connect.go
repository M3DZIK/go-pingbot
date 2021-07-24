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

func Connect(retry ...int8) {
	if len(retry) == 0 {
		retry = append(retry, 0)
	}

	if retry[0] == 2 {
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Mongo_URI))
	if common.CheckErr(err, "connect to db") {
		time.Sleep(2 * time.Second)
		Connect(retry[0] + 1)

		return
	}

	err = Client.Ping(ctx, readpref.Primary())
	if common.CheckErr(err, "ping db") {
		time.Sleep(2 * time.Second)
		Connect(retry[0] + 1)

		return
	}

	Coll = Client.Database(config.Mongo_DB).Collection(config.Mongo_Collection)
}
