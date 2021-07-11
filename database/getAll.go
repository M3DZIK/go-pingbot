package database

import (
	"context"
	"time"

	"gitlab.com/gaming0skar123/go/pingbot/config"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAll() ([]URL, error) {
	collection := DB.Collection(config.Mongo_Collection)

	var results []URL

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
