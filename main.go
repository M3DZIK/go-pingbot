package main

import (
	"gitlab.com/gaming0skar123/go/pingbot/backend"
	"gitlab.com/gaming0skar123/go/pingbot/common"
	"gitlab.com/gaming0skar123/go/pingbot/config"
	"gitlab.com/gaming0skar123/go/pingbot/database"
	"gitlab.com/gaming0skar123/go/pingbot/website"
)

var log = common.Log

func main() {
	log.Info("You're using verion: ", config.Version)

	database.Connect()

	go website.Server()
	backend.Run()
}
