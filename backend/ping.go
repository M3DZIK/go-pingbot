package backend

import (
	"net/http"

	"gitlab.com/gaming0skar123/go/pingbot/common"
	"gitlab.com/gaming0skar123/go/pingbot/database"
)

var AmountSuccess uint8
var AmountErr uint8

var checkErr = common.CheckErr

func ping() {
	results, err := database.GetAll()
	if checkErr(err, "Get All res from DB") {
		return
	}

	for _, value := range results {
		go loop(value)
	}
}

func loop(value database.URL) {
	r, err := http.Get(value.URL)
	if checkErr(err, "Ping URL") {
		AmountErr++
		return
	}

	if r.StatusCode >= 200 && r.StatusCode < 400 {
		AmountSuccess++
	} else {
		AmountErr++
	}
}
