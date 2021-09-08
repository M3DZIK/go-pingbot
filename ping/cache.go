package ping

import (
	"time"

	"github.com/MedzikUser/go-utils/common"
	"gitlab.com/gaming0skar123/go/pingbot/database/mongo"
)

var CacheURL []string

func Cache() {
	CacheURL = CacheURL[:0]

	err := common.Retry(15, 1*time.Second, func() error {
		results, err := mongo.GetAll()
		if err != nil {
			return err
		}

		for _, value := range results {
			CacheURL = append(CacheURL, value.URL)
		}

		return nil
	})

	if err != nil {
		log.Error(err)
	}
}
