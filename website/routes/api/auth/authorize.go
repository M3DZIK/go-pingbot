package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authorize(c *gin.Context) (bool, jwt.MapClaims, error) {
	const BEARER_SCHEMA = "Bearer"

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return false, nil, nil
	}
	tokenString := authHeader[len(BEARER_SCHEMA):]
	token, claims, err := ValidateToken(tokenString)
	if err != nil {
		return false, claims, err
	}

	return token.Valid, claims, err
}
