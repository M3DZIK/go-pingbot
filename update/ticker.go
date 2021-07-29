package update

import (
	"time"

	"gitlab.com/gaming0skar123/go/pingbot/config"
)

func Ticker() {
	// Check on start
	Update()

	ticker := time.NewTicker(config.Toml.AutoUpdate.Check * time.Minute)

	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			Update()

		case <-quit:
			ticker.Stop()

			return
		}
	}
}
