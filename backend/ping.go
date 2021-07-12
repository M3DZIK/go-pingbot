package backend

import (
	"net/http"

	"github.com/tcnksm/go-httpstat"
	"gitlab.com/gaming0skar123/go/pingbot/common"
	"gitlab.com/gaming0skar123/go/pingbot/database"
)

var AmountSuccess uint
var AmountErr uint

var checkErr = common.CheckErr

func ping() {
	results, err := database.GetAll()
	if checkErr(err, "get all from db") {
		return
	}

	for _, value := range results {
		go loop(value)
	}
}

func loop(value database.URL) {
	req, err := http.NewRequest("GET", value.URL, nil)
	if checkErr(err, "new http request") {
		AmountErr++
		return
	}

	var result httpstat.Result

	ctx := httpstat.WithHTTPStat(req.Context(), &result)
	req = req.WithContext(ctx)

	client := http.DefaultClient
	r, err := client.Do(req)
	if checkErr(err, "ping url") {
		AmountErr++
		return
	}

	if r.StatusCode >= 200 && r.StatusCode < 400 {
		AmountSuccess++
	} else {
		AmountErr++
	}
}
