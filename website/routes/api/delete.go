package api

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/gaming0skar123/go/pingbot/database"
)

func Delete(c *gin.Context) {
	var post URLType
	c.BindJSON(&post)

	url := c.Param("url")

	d, err := base64.StdEncoding.DecodeString(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error Parsing Base64!",
		})

		return
	}

	url = string(d)

	r, err := database.Delete(&database.URL{
		URL: url,
	})

	if r.DeletedCount <= 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Not found!",
		})

		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error Deleting from Database!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Deleted",
	})
}
