package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/gaming0skar123/go/pingbot/backend"
)

func GetAll(c *gin.Context) {
	backend.Cache()

	c.JSON(http.StatusOK, json{
		"success": true,
		"db":      backend.CacheURL,
	})
}
