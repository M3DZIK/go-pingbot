package mongo

import (
	"context"
	"time"

	"github.com/MedzikUser/go-utils/common"
	"gitlab.com/gaming0skar123/go/pingbot/config"
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
		if common.CheckErr(err, "connect to db") {
			time.Sleep(2 * time.Second)

			return err
		}

		err = Client.Ping(ctx, readpref.Primary())
		if common.CheckErr(err, "ping db") {
			time.Sleep(2 * time.Second)

			return err
		}

		Coll = Client.Database(config.Mongo_DB).Collection(config.Mongo_Collection)

		return nil
	})
}
