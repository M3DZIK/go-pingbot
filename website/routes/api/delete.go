package api

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/gaming0skar123/go/pingbot/config"
	"gitlab.com/gaming0skar123/go/pingbot/database/mongo"
)

func Delete(c *gin.Context) {
	const BEARER_SCHEMA = "Password"
	authHeader := c.GetHeader("Authorization")
	passwordString := authHeader[len(BEARER_SCHEMA)+1:]

	if passwordString != config.Password {
		c.JSON(http.StatusUnauthorized, json{
			"success": false,
			"message": "Unauth!",
		})

		return
	}

	url := c.Param("url")

	d, err := base64.StdEncoding.DecodeString(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, json{
			"success": false,
			"message": "Error Parsing Base64!",
		})

		return
	}

	url = string(d)

	r, err := mongo.Delete(url)

	if r.DeletedCount <= 0 {
		c.JSON(http.StatusNotFound, json{
			"success": false,
			"message": "Not Found!",
		})

		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, json{
			"success": false,
			"message": "Error Deleting from Database!",
		})

		return
	}

	c.JSON(http.StatusOK, json{
		"success": true,
		"message": "Deleted!",
	})
}
