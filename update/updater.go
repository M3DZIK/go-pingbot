package update

import (
	"time"

	"gitlab.com/gaming0skar123/go/pingbot/config"
)

func Updater() {
	// Check on start
	CheckUpdate()

	ticker := time.NewTicker(config.Latest_Version_Check)

	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			CheckUpdate()

		case <-quit:
			ticker.Stop()

			return
		}
	}
}
