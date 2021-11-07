package api

import (
	"net/http"
	"runtime"

	"github.com/MedzikUser/go-utils/common"
	"github.com/MedzikUser/go-utils/stats"
	"github.com/gin-gonic/gin"
	"gitlab.com/MedzikUser/go/pingbot/config"
	"gitlab.com/MedzikUser/go/pingbot/ping"
)

func Status(c *gin.Context) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	mem := stats.Memory()
	cpu, err := stats.CPU()
	if err != nil {
			common.Log.Error("cpu stat", err)
	}

	var p json

	if config.Toml.Backend.Enabled {
		p = json{
			"all":     ping.Status.Error + ping.Status.Success,
			"success": ping.Status.Success,
			"error":   ping.Status.Error,
		}
	} else {
		p = nil
	}

	c.JSON(http.StatusOK, json{
		"ping": p,
		"sys": json{
			"uptime": common.Uptime(config.StartTime),
			"pid":    cpu.PID,
			"os":     runtime.GOOS,
			"memory": json{
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
	})
}
