package backend

import (
	"time"

	"github.com/MedzikUser/go-utils/common"
	"gitlab.com/gaming0skar123/go/pingbot/database/mongo"
)

var cacheURL []string

func cache() {
	cacheURL = cacheURL[:0]

	err := common.Retry(1, 1*time.Second, func() error {
		results, err := mongo.GetAll()
		if err != nil {
			return err
		}

		for _, value := range results {
			cacheURL = append(cacheURL, value.URL)
		}

		return nil
	})

	if err != nil {
		log.Error(err)
	}
}
