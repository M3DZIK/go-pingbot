package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func Delete(url *URL) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := Coll.DeleteOne(ctx, url)

	return result, err
}
