package website

import (
	"os"
	"strconv"

	"github.com/MedzikUser/go-utils/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/medzikuser/go-pingbot/config"
	"github.com/medzikuser/go-pingbot/website/routes"
	"github.com/medzikuser/go-pingbot/website/routes/api"
)

var router *gin.Engine

func Server() {
	// Disable GIN Debug
	gin.SetMode(gin.ReleaseMode)

	router = gin.New()

	// Fix cors
	configCors := cors.DefaultConfig()
	configCors.AllowOrigins = []string{"https://www.pingbot.cf"}

	router.Use(cors.New(configCors))

	router.GET("/", routes.Index)

	api.ApplyRoutes(router)

	err := router.Run(":" + strconv.Itoa(config.Toml.HTTP.Port))
	if err != nil {
		common.Log.Error("gin start http server", err)

		os.Exit(1)
	}
}
