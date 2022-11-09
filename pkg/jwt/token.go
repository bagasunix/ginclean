package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(jwtKey string, claims Claims) (token string, err error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(jwtKey))
}
