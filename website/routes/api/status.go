package api

import (
	"net/http"
	"os"
	"runtime"

	"github.com/MedzikUser/go-utils/common"
	"github.com/MedzikUser/go-utils/stats"
	"github.com/gin-gonic/gin"
	"gitlab.com/gaming0skar123/go/pingbot/backend"
	"gitlab.com/gaming0skar123/go/pingbot/config"
)

func Status(c *gin.Context) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	mem := stats.Memory()
	cpu, err := stats.CPU()
	common.CheckErr(err, "cpu stat")

	var ping json

	if config.Toml.Backend.Enabled {
		ping = json{
			"all":     backend.Status.Error + backend.Status.Success,
			"success": backend.Status.Success,
			"err":     backend.Status.Error,
		}
	} else {
		ping = nil
	}

	c.JSON(http.StatusOK, json{
		"ping": ping,
		"sys": json{
			"pid": os.Getpid(),
			"os":  runtime.GOOS,
			"mem": json{
				"alloc":      mem.Alloc,
				"totalalloc": mem.TotalAlloc,
				"sys":        mem.Sys,
				"numgc":      mem.NumGC,
			},
			"cpu": json{
				"usage": cpu.Usage,
				"num":   cpu.Num,
				"arch":  cpu.Arch,
			},
		},
		"v": json{
			"go":      runtime.Version(),
			"release": config.Version,
			"build":   config.Build,
		},
		"node": json{
			"cluster": config.Toml.Cluster.ID,
			"node":    config.Toml.Cluster.Node,
			"uptime":  common.Uptime(config.StartTime),
		},
	})
}
