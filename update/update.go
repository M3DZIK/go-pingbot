package update

import (
	"net/http"
	"os"

	"github.com/inconshreveable/go-update"
	"gitlab.com/gaming0skar123/go/pingbot/common"
	"gitlab.com/gaming0skar123/go/pingbot/config"
)

var log = common.Log

func Update() {
	log.Warn("Updating...")

	url := config.Latest_Binary

	resp, err := http.Get(url)
	if common.CheckErr(err, "downloading latest binary") {
		return
	}

	defer resp.Body.Close()

	err = update.Apply(resp.Body, update.Options{})
	if common.CheckErr(err, "self-update binary") {
		return
	}

	log.Info("Updated!")

	defer os.Exit(0)
}
