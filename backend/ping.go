package backend

import (
	"net/http"

	"gitlab.com/gaming0skar123/go/pingbot/common"
	"gitlab.com/gaming0skar123/go/pingbot/database/mongo"
)

var checkErr = common.CheckErr

func ping() {
	results, err := mongo.GetAll()
	if checkErr(err, "get all from db") {
		return
	}

	for _, value := range results {
		go loop(value)
	}
}

func loop(value mongo.URL) {
	req, err := http.NewRequest("GET", value.URL, nil)
	if checkErr(err, "new http request") {
		Status.Error++
		return
	}

	client := http.DefaultClient
	r, err := client.Do(req)
	if checkErr(err, "ping url") {
		Status.Error++
		return
	}

	if r.StatusCode >= 200 && r.StatusCode < 400 {
		Status.Success++
	} else {
		Status.Error++
	}
}
