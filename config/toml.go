package config

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"gitlab.com/gaming0skar123/go/pingbot/common"
)

type tomlConfig struct {
	HTTP       httpConfig
	Backend    backendConfig
	AutoUpdate autoUpdateConfig
	Cluster    clusterConfig
}

type httpConfig struct {
	Enabled bool
	Port    int
}

type backendConfig struct {
	Enabled bool
	Ping    time.Duration
}

type autoUpdateConfig struct {
	Enabled bool
	Check   time.Duration
}

type clusterConfig struct {
	ID   int
	Node int
}

var Toml tomlConfig

func init() {
	_, err := toml.DecodeFile("./config.toml", &Toml)

	if common.CheckErr(err, "decode toml config") {
		if err.Error() == "open ./config.toml: no such file or directory" {
			err := DownloadFile("config.toml", "https://github.com/"+GH_Repo+"/raw/main/config.schema.toml")

			if !common.CheckErr(err, "download default config") {
				_, err = toml.DecodeFile("./config.toml", &Toml)
				common.CheckErr(err, "decode toml config")
			}
		}
	}
}

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
