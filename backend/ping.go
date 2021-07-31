package backend

import (
	"context"
	"net/http"
	"time"

	"gitlab.com/gaming0skar123/go/pingbot/common"
	"gitlab.com/gaming0skar123/go/pingbot/database/mongo"
)

var checkErr = common.CheckErr

func ping() {
	results, err := mongo.GetAll()
	if checkErr(err, "get keys from db") {
		return
	}

	for _, value := range results {
		go loop(value)
	}
}

func loop(value mongo.URL) {
	// Timeout 1 minute
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", value.URL, nil)
	if checkErr(err, "new http request") {
		Status.Error++
		return
	}

	client := http.DefaultClient
	r, err := client.Do(req)
	if err != nil {
		Status.Error++
		return
	}

	if r.StatusCode >= 200 && r.StatusCode < 400 {
		Status.Success++
	} else {
		Status.Error++
	}
}
