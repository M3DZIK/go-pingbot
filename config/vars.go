package config

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const GH_Repo = "MedzikUser/go-pingbot"

var (
	Version   = "dev"
	Build     = ""
	StartTime time.Time
	UserAgent = fmt.Sprintf("PingBot/%s go/%s", Version, strings.Replace(runtime.Version(), "go", "", 1))
)
