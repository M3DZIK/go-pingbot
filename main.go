package main

import (
	"github.com/jpillora/opts"
	"gitlab.com/gaming0skar123/go/pingbot/backend"
	"gitlab.com/gaming0skar123/go/pingbot/common"
	"gitlab.com/gaming0skar123/go/pingbot/config"
	_ "gitlab.com/gaming0skar123/go/pingbot/database"
	"gitlab.com/gaming0skar123/go/pingbot/update"
	"gitlab.com/gaming0skar123/go/pingbot/website"
)

type CMDOptions struct {
	Update bool `opts:"help=automatic updates"`
}

var log = common.Log

func main() {
	cmd := CMDOptions{
		Update: true,
	}

	opts.Parse(&cmd)

	log.Info("You're using verion: ", config.Version)

	if cmd.Update {
		go update.Updater()
	} else {
		log.Warn("Auto Update -> Disabled")
	}

	go website.Server()
	backend.Run()
}
