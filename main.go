package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/MedzikUser/go-utils/common"
	"github.com/MedzikUser/go-utils/updater"
	"gitlab.com/gaming0skar123/go/pingbot/backend"
	"gitlab.com/gaming0skar123/go/pingbot/config"
	"gitlab.com/gaming0skar123/go/pingbot/database/mongo"
	"gitlab.com/gaming0skar123/go/pingbot/website"
)

var log = common.Log

func main() {
	log.Info("You're using verion: ", config.Version)

	var wg sync.WaitGroup

	mongo.Connect()

	if config.Toml.AutoUpdate.Enabled {
		wg.Add(1)

		client := updater.Client{
			GitHub:      config.GH_Repo,
			GitHubToken: config.GH_Token,
			Version:     config.Version,
			Binary:      "pingbot.out",
			CheckEvery:  config.Toml.AutoUpdate.Check * time.Minute,
			AfterUpdate: func() {
				log.Info("Updated!")

				if !config.Toml.Options.Stop_After_Ping {
					os.Exit(0)
				}
			},
			Major: false,
		}

		go client.AutoUpdater()
	} else {
		log.Warn("Auto Update -> Disabled")
	}

	if config.Toml.Options.Stop_After_Ping {
		dbNum := backend.StopAfterPing()

		fmt.Println()

		log.Info("DB Size -> ", dbNum)
		log.Info("Pinged  -> ", backend.Status.Success+backend.Status.Error)
		log.Info("Success -> ", backend.Status.Success)
		log.Info("Error   -> ", backend.Status.Error)

		os.Exit(0)
	}

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

	config.StartTime = time.Now()

	wg.Wait()
}
