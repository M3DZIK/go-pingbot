package auth

import (
	"github.com/golang-jwt/jwt"
	"gitlab.com/gaming0skar123/go/pingbot/config"
)

func ValidateToken(encodedToken string) (*jwt.Token, jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(encodedToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_Secret), nil
	})

	return token, claims, err
}
