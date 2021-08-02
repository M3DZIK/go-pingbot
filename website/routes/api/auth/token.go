package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"gitlab.com/gaming0skar123/go/pingbot/config"
)

type authCustomClaims struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) string {
	claims := &authCustomClaims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    username,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(config.JWT_Secret))
	if err != nil {
		panic(err)
	}

	return t
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}
		return []byte(config.JWT_Secret), nil
	})
}
