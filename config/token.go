package config

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(username string) (tokenString string, err error) {
	expirarionTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirarionTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(os.Getenv("JWT_SECRET_KEY"))
	return
}
