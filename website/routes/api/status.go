package api

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"runtime"

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

	c.JSON(http.StatusOK, json{
		"ping": json{
			"all":     backend.Status.Error + backend.Status.Success,
			"success": backend.Status.Success,
			"err":     backend.Status.Error,
		},
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
		},
	})
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
