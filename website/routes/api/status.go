package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/gaming0skar123/go/pingbot/backend"
)

func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"ping": gin.H{
			"success": backend.AmountSuccess,
			"err":     backend.AmountErr,
		},
	})
}
