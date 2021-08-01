package api

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/struCoder/pidusage"
	"gitlab.com/gaming0skar123/go/pingbot/backend"
	"gitlab.com/gaming0skar123/go/pingbot/common"
	"gitlab.com/gaming0skar123/go/pingbot/config"
)

func Status(c *gin.Context) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	pid := os.Getpid()

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
			"pid": pid,
			"os":  runtime.GOOS,
			"mem": json{
				"alloc":      mb(m.Alloc),
				"totalalloc": mb(m.TotalAlloc),
				"sys":        mb(m.Sys),
				"numgc":      m.NumGC,
			},
			"cpu": json{
				"usage": cpu(pid),
				"num":   runtime.NumCPU(),
				"arch":  runtime.GOARCH,
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
			//"uptime":  time.Since(config.StartTime).String(),
			"uptime": uptime(),
		},
	})
}

func uptime() string {
	t := time.Since(config.StartTime)

	var uptime string

	var (
		y int
		d int
	)

	h := round(t.Hours())
	m := round(t.Minutes())
	s := round(t.Seconds())

	for h/24 > 0 {
		d++
		h -= 24
	}

	for d/365 > 0 {
		y++
		d -= 365
	}

	if y > 0 {
		uptime += strconv.Itoa(y) + "y "
	}

	if d > 0 {
		uptime += strconv.Itoa(d) + "d "
	}

	if h > 0 {
		uptime += strconv.Itoa(h) + "h "
	}

	if m > 0 {
		uptime += strconv.Itoa(m-(round(t.Hours())*60)) + "m "
	}

	if s > 0 {
		uptime += strconv.Itoa(s-(m*60)) + "s"
	}

	return uptime
}

func round(val float64) int {
	if val < 0 {
		return int(val - 1.0)
	}

	return int(val)
}

func mb(b uint64) string {
	return fmt.Sprintf("%d MB", b/1000/1000)
}

func cpu(pid int) *string {
	sysInfo, err := pidusage.GetStat(pid)
	if common.CheckErr(err, "get cpu stat") {
		return nil
	}

	s := fmt.Sprint(math.Round(sysInfo.CPU*100)/100, "%")

	return &s
}
