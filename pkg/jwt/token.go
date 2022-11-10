package jwt

import (
	"errors"
	"fmt"
	"time"

	envs "github.com/bagasunix/ginclean/pkg/env"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(jwtKey string, claims Claims) (token string, err error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return t.SignedString([]byte(jwtKey))
}

func ValidateToken(signedToken string) (err error) {
	conf, _ := envs.LoadEnv()
	t, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]))
		}
		return []byte(conf.JwtSecret), nil
	})

	if err != nil {
		return
	}

	claims, ok := t.Claims.(*Claims)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	return err
}
