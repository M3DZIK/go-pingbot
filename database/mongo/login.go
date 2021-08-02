package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func Login(username, password string) (*mongo.SingleResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r := Coll.FindOne(ctx, json{
		"_id":      username,
		"password": password,
	})

	return r, r.Err()
}
