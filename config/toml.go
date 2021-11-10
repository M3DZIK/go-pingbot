package config

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/MedzikUser/go-utils/common"
)

type tomlConfig struct {
	HTTP       httpConfig
	Backend    backendConfig
	AutoUpdate autoUpdateConfig
	MongoDB    mongoDBConfig
	Options    optionsConfig
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

type mongoDBConfig struct {
	Database   string
	Collection string
}

type optionsConfig struct {
	Stop_After_Ping         bool
	Stop_After_Ping_Timeout time.Duration
}

var Toml tomlConfig

func init() {
	_, err := toml.DecodeFile("./config.toml", &Toml)

	if err != nil {
		common.Log.Error("decode toml config: ", err)
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
