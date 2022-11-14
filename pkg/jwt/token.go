package jwt

import (
	"fmt"
	"time"

	envs "github.com/bagasunix/ginclean/pkg/env"
	"github.com/bagasunix/ginclean/pkg/errors"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

func GenerateToken(jwtKey string, claims Claims) (token string, err error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return t.SignedString([]byte(jwtKey))
}

func ValidateToken(logs *zap.Logger, signedToken string) (claims *Claims, err error) {
	conf, _ := envs.LoadEnv()
	t, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err = errors.CustomError(fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]))
			return nil, err
		}
		return []byte(conf.JwtSecret), nil
	})

	if err != nil {
		return nil, errors.ErrSomethingWrong(logs, err)
	}

	claims, ok := t.Claims.(*Claims)
	if !ok {
		err = errors.CustomError("couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.CustomError("token expired")
		return
	}

	return claims, err
}

func ValidateRefreshToken(logs *zap.Logger, signedToken string) (claims *Claims, err error) {
	conf, _ := envs.LoadEnv()
	t, _ := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return claims, errors.ErrSomethingWrong(logs, errors.CustomError(fmt.Sprintf("unexpected signing method: %v", token.Header["alg"])))
		}
		return []byte(conf.JwtSecretRefresh), nil
	})

	claims, ok := t.Claims.(*Claims)
	if !ok {
		return claims, errors.ErrSomethingWrong(logs, errors.CustomError("couldn't parse claims"))
	}

	return claims, err
}
