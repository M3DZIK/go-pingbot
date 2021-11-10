package mongo

import (
	"context"
	"time"

	"github.com/MedzikUser/go-utils/common"
	"github.com/medzikuser/go-pingbot/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client mongo.Client
var Coll *mongo.Collection

func Connect() error {
	return common.Retry(5, 2*time.Second, func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		Client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Mongo_URI))
		if err != nil {
			common.Log.Error("connect to db", err)

			time.Sleep(2 * time.Second)

			return err
		}

		err = Client.Ping(ctx, readpref.Primary())
		if err != nil {
			common.Log.Error("ping db", err)

			time.Sleep(2 * time.Second)

			return err
		}

		Coll = Client.Database(config.Toml.MongoDB.Database).Collection(config.Toml.MongoDB.Collection)

		return nil
	})
}
