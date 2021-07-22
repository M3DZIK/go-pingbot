package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func Insert(url *URL) (*mongo.InsertOneResult, error) {ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := Coll.InsertOne(ctx, url)

	return result, err
}
