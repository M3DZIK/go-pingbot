package ping

import (
	"time"

	"gitlab.com/gaming0skar123/go/pingbot/config"
)

func Ticker() {
	// On Start
	Cache()
	ping()

	ticker := time.NewTicker(config.Toml.Backend.Ping * time.Minute)

	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			ping()
		case <-quit:
			ticker.Stop()

			return
		}
	}
}
