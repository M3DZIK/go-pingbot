package backend

import (
	"time"

	"gitlab.com/gaming0skar123/go/pingbot/config"
	"gitlab.com/gaming0skar123/go/pingbot/database/redis"
)

func Ticker() {
	// Ping on Start
	redis.StatusGet()
	ping()
	redis.StatusGet()

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
