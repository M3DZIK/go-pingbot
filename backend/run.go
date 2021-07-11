package backend

import (
	"time"
)

func Run() {
	ping()

	ticker := time.NewTicker(2 * time.Minute)

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
