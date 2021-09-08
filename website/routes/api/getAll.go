package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/gaming0skar123/go/pingbot/ping"
)

func GetAll(c *gin.Context) {
	ping.Cache()

	c.JSON(http.StatusOK, json{
		"success": true,
		"db":      ping.CacheURL,
	})
}
