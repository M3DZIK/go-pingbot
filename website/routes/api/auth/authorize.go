package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthorizeJWT(c *gin.Context) bool {
	const BEARER_SCHEMA = "Bearer"
	authHeader := c.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA):]
	token, err := ValidateToken(tokenString)
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		fmt.Println(claims)
		return true
	} else {
		fmt.Println(err)
		return false
	}
}
