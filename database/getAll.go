package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAll() ([]URL, error) {
	var results []URL

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := Coll.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
