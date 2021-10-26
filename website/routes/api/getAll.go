package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/MedzikUser/go/pingbot/ping"
)

func GetAll(c *gin.Context) {
	ping.Cache()

	c.JSON(http.StatusOK, json{
		"db": ping.CacheURL,
	})
}
