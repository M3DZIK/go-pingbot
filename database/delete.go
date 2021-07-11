package database

import (
	"context"
	"time"

	"gitlab.com/gaming0skar123/go/pingbot/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func Delete(url *URL) (*mongo.DeleteResult, error) {
	collection := DB.Collection(config.Mongo_Collection)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, url)

	return result, err
}
