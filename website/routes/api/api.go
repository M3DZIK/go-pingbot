package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/gaming0skar123/go/pingbot/website/routes/api/auth"
)

func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/url", GetAll)
		api.POST("/url", Insert)
		api.DELETE("/url/:url", Delete)

		api.POST("/login", auth.Login)

		api.GET("/status", Status)
	}
}
