package main

import (
	"os"
	"os/signal"
	"syscall"

	"gitlab.com/gaming0skar123/go/pingbot/backend"
	"gitlab.com/gaming0skar123/go/pingbot/common"
	"gitlab.com/gaming0skar123/go/pingbot/config"
	"gitlab.com/gaming0skar123/go/pingbot/database/mongo"
	"gitlab.com/gaming0skar123/go/pingbot/database/redis"
	"gitlab.com/gaming0skar123/go/pingbot/update"
	"gitlab.com/gaming0skar123/go/pingbot/website"
)

var log = common.Log

func main() {
	log.Info("You're using verion: ", config.Version)

	mongo.Connect()
	redis.Connect()

	if config.Toml.HTTP.Enabled {
		go website.Server()
	} else {
		log.Warn("HTTP Server -> Disabled")
	}

	if config.Toml.Backend.Enabled {
		go backend.Ticker()
	} else {
		log.Warn("Backend -> Disabled")
	}

	if config.Toml.AutoUpdate.Enabled {
		go update.Ticker()
	} else {
		log.Warn("Auto Update -> Disabled")
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	func() {
		<-c
		log.Warn("Stopping...")
		os.Exit(1)
	}()
}
