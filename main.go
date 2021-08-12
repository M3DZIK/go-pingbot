package main

import (
	"sync"
	"time"

	"github.com/MedzikUser/go-utils/common"
	"gitlab.com/gaming0skar123/go/pingbot/backend"
	"gitlab.com/gaming0skar123/go/pingbot/config"
	"gitlab.com/gaming0skar123/go/pingbot/database/mongo"
	"gitlab.com/gaming0skar123/go/pingbot/update"
	"gitlab.com/gaming0skar123/go/pingbot/website"
)

var log = common.Log

func main() {
	log.Info("You're using verion: ", config.Version)

	var wg sync.WaitGroup

	mongo.Connect()

	if config.Toml.HTTP.Enabled {
		wg.Add(1)
		go website.Server()
	} else {
		log.Warn("HTTP Server -> Disabled")
	}

	if config.Toml.Backend.Enabled {
		wg.Add(1)
		go backend.Ticker()
	} else {
		log.Warn("Backend -> Disabled")
	}

	if config.Toml.AutoUpdate.Enabled {
		wg.Add(1)
		go update.Ticker()
	} else {
		log.Warn("Auto Update -> Disabled")
	}

	config.StartTime = time.Now()

	wg.Wait()
}
