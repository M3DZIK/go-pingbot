package api

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"gitlab.com/gaming0skar123/go/pingbot/backend"
	"gitlab.com/gaming0skar123/go/pingbot/config"
)

func Status(c *gin.Context) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"ping": gin.H{
			"all":     backend.AmountSuccess + backend.AmountErr,
			"success": backend.AmountSuccess,
			"err":     backend.AmountErr,
		},
		"mem": gin.H{
			"alloc":      MB(m.Alloc),
			"totalalloc": MB(m.TotalAlloc),
			"sys":        MB(m.Sys),
			"numgc":      m.NumGC,
		},
		"v": config.Version,
	})
}

func MB(b uint64) string {
	return fmt.Sprintf("%d MB", b/1000/1000)
}
