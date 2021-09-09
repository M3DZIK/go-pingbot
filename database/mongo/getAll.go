package mongo

import (
	"context"
	"time"
)

func GetAll() ([]URL, error) {
	var results []URL

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := Coll.Find(ctx, URL{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
