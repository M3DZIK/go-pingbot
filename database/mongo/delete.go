package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Delete(url, owner string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := Coll.DeleteOne(ctx, bson.M{"_id": url, "owner": owner})
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return r, err
}
