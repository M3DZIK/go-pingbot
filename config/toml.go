package config

import (
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

	common.CheckErr(err, "decode toml config")
}
