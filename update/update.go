package update

import (
	"os"

	"github.com/MedzikUser/go-github-selfupdate/selfupdate"
	"github.com/MedzikUser/go-utils/common"
	"github.com/blang/semver/v4"
	"gitlab.com/gaming0skar123/go/pingbot/config"
)

var log = common.Log

func Update() {
	repo := config.GH_Repo

	latest, found, err := selfupdate.DetectLatest(repo)
	if common.CheckErr(err, "detecting letest version") {
		return
	}

	v, err := semver.Parse(config.Version)
	if common.CheckErr(err, "parse version") || !found || latest.Version.LTE(v) {
		return
	}

	log.Warn("Updating...")

	exe, err := os.Executable()
	if common.CheckErr(err, "locate executable path") {
		return
	}

	err = selfupdate.UpdateTo(latest.AssetURL, exe, "pingbot.out")
	if common.CheckErr(err, "update binary") {
		return
	}

	log.Info("Updated!")

	defer os.Exit(0)
}
