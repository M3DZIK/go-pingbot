package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/MedzikUser/go-utils/common"
	"github.com/MedzikUser/go-utils/updater"
	"github.com/jpillora/opts"
	"gitlab.com/gaming0skar123/go/pingbot/backend"
	"gitlab.com/gaming0skar123/go/pingbot/config"
	"gitlab.com/gaming0skar123/go/pingbot/database/mongo"
	"gitlab.com/gaming0skar123/go/pingbot/website"
)

var log = common.Log

type cmdOpts struct {
	Update bool `opts:"help=update version to latest e.g. if update is major"`
}

func main() {
	log.Info("You're using verion: ", config.Version)

	var wg sync.WaitGroup

	mongo.Connect()

	c := cmdOpts{}

	opts.Parse(&c)

	if c.Update {
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

		err := client.Update()
		if err != nil && err.Error() == "major update" {
			fmt.Print("Update to new major version? (y/N) ")
			reader := bufio.NewReader((os.Stdin))
			char, _, err := reader.ReadRune()
			if err != nil {
				fmt.Println(err)
			}

			switch char {
			case 'y':
				client.Major = true
				err := client.Update()
				if err != nil {
					log.Error(err)
					os.Exit(1)
				}

			case 'Y':
				client.Major = true
				err := client.Update()
				if err != nil {
					log.Error(err)
					os.Exit(1)
				}

			default:
				log.Warn("Canceled!")
				os.Exit(2)
			}
		} else if err != nil {
			log.Error(err)
			os.Exit(1)
		} else {
			log.Info("You're using latest version!")
			os.Exit(0)
		}
	}

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
		backend.StopAfterPing()

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
