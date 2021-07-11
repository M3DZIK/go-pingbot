package database

import (
	"context"
	"time"

	"gitlab.com/gaming0skar123/go/pingbot/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func Insert(url *URL) (*mongo.InsertOneResult, error) {
	collection := DB.Collection(config.Mongo_Collection)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, url)

	return result, err
}
