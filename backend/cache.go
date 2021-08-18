package backend

import (
	"time"

	"github.com/MedzikUser/go-utils/common"
	"gitlab.com/gaming0skar123/go/pingbot/database/mongo"
)

var cacheURL []string

func cache(retry int) {
	cacheURL = cacheURL[:0]

	results, err := mongo.GetAll()
	if common.CheckErr(err, "get documents from db") {
		if retry == 5 {
			time.Sleep(500 * time.Millisecond)
			cache(retry + 1)
		} else {
			return
		}
	}

	for _, value := range results {
		cacheURL = append(cacheURL, value.URL)
	}
}
