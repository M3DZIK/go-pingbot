package backend

import (
	"gitlab.com/gaming0skar123/go/pingbot/database/mongo"
)

var cacheURL []string

func cache() {
	cacheURL = cacheURL[:0]

	results, err := mongo.GetAll()
	if checkErr(err, "get keys from db") {
		return
	}

	for _, value := range results {
		cacheURL = append(cacheURL, value.URL)
	}
}
