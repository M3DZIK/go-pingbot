package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func Delete(url string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := Coll.DeleteOne(ctx, json{
		"url": url,
	})

	return result, err
}
