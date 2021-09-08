package ping

import (
	"os"
	"time"

	"github.com/MedzikUser/go-utils/common"
	"gitlab.com/gaming0skar123/go/pingbot/config"
)

var log = common.Log

func StopAfterPing() {
	Cache()
	num := ping()

	// timeout
	go func() {
		timeout := config.Toml.Options.Stop_After_Ping_Timeout
		if timeout == 0 {
			return
		}

		time.Sleep(timeout * time.Second)

		log.Warn("DB Size -> ", num)
		log.Warn("Pinged  -> ", Status.Success+Status.Error)
		log.Warn("Success -> ", Status.Success)
		log.Warn("Error   -> ", Status.Error)

		os.Exit(1)
	}()

	for {
		if int64(num) == Status.Success+Status.Error {
			log.Info("DB Size -> ", num)
			log.Info("Pinged  -> ", Status.Success+Status.Error)
			log.Info("Success -> ", Status.Success)
			log.Info("Error   -> ", Status.Error)

			break
		}
	}
}
