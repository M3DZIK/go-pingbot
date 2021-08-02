package api

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/gaming0skar123/go/pingbot/database/mongo"
	"gitlab.com/gaming0skar123/go/pingbot/website/routes/api/auth"
)

func Delete(c *gin.Context) {
	valid, claims, err := auth.Authorize(c)
	if !valid || err != nil {
		c.JSON(http.StatusBadRequest, json{
			"success": false,
			"message": "Unauthed!",
		})

		return
	}

	var user string

	for key, val := range claims {
		if key == "user" {
			user = fmt.Sprintf("%q", val)
			break
		}
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

	r, err := mongo.Delete(url, user)

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
