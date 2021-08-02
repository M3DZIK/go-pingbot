package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func Register(username, password string) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := Coll.InsertOne(ctx, json{
		"_id": username,
		"password": password,
	})

	return result, err
}
