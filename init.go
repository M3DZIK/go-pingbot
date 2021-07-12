package main

import (
	"github.com/jpillora/opts"
	"gitlab.com/gaming0skar123/go/pingbot/update"
)

type CMDOptions struct {
	Update bool `opts:"help=automatic updates"`
}

func init() {
	cmd := CMDOptions{
		Update: true,
	}

	opts.Parse(&cmd)

	if cmd.Update {
		go update.Updater()
	} else {
		log.Warn("Auto Update -> Disabled")
	}
}
