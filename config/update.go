package config

import (
	"fmt"
	"runtime"
	"time"
)

var Latest_Version_API = "https://gitlab.com/gaming0skar123/go/pingbot/-/jobs/artifacts/main/raw/VERSION?job=build"
var Latest_Binary = fmt.Sprintf("https://gitlab.com/gaming0skar123/go/pingbot/-/jobs/artifacts/main/raw/pingbot-%s-%s?job=build", runtime.GOOS, runtime.GOARCH)
var Latest_Version_Check = 2 * time.Minute
