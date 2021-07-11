package update

import (
	"io/ioutil"
	"net/http"

	"github.com/blang/semver/v4"
	"gitlab.com/gaming0skar123/go/pingbot/common"
	"gitlab.com/gaming0skar123/go/pingbot/config"
)

func CheckUpdate() {
	url := config.Latest_Version_API

	resp, err := http.Get(url)
	if common.CheckErr(err, "GET latest version from API") {
		return
	}

	defer resp.Body.Close()

	r := resp.Body

	b, _ := ioutil.ReadAll(r)

	v := string(b)

	confVer, err := semver.Make(config.Version)
	if common.CheckErr(err, "make version from config") {
		return
	}

	apiVer, err := semver.Make(v)
	if common.CheckErr(err, "make version from API") {
		return
	}

	ver := confVer.Compare(apiVer)

	if ver == -1 {
		Update()
	}
}
