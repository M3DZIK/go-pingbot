package auth

import (
	"crypto/sha512"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gitlab.com/gaming0skar123/go/pingbot/config"
	"gitlab.com/gaming0skar123/go/pingbot/database/mongo"
)

func Login(c *gin.Context) {
	var post UserAuth
	err := c.BindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, json{
			"success": false,
			"message": "Error Binding JSON!",
		})

		return
	}

	if post.Username == "" || post.Password == "" {
		c.JSON(http.StatusBadRequest, json{
			"success": false,
			"message": "Invalid POST param!",
		})

		return
	}

	if len(post.Username) < 6 || len(post.Password) < 8 {
		c.JSON(http.StatusBadRequest, json{
			"success": false,
			"message": "Username or Password is too short!",
		})

		return
	}

	h := sha512.Sum512([]byte(post.Password))
	s := fmt.Sprintf("%x", h[:])

	_, err = mongo.Login(post.Username, s)
	if err != nil {
		_, err = mongo.Register(post.Username, s)
		if err != nil {
			c.JSON(http.StatusNotFound, json{
				"success": false,
				"message": "Can't register!",
			})
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": post.Username,
		"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":  time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.JWT_Secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, json{
			"success": false,
			"message": "Error sign token!",
		})

		return
	}

	c.JSON(http.StatusOK, json{
		"success": true,
		"token":   tokenString,
	})
}
