package jwt

import (
	"time"

	"github.com/bagasunix/ginclean/server/domains/entities"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	User *entities.Account `json:"account,omitempty"`
	jwt.StandardClaims
}

type ClaimsBuilder struct {
	claims *Claims
}

func (c *ClaimsBuilder) User(user *entities.Account) *ClaimsBuilder {
	c.claims.User = user
	return c
}

func (c *ClaimsBuilder) ExpiresAt(expiresAt time.Time) *ClaimsBuilder {
	c.claims.ExpiresAt = expiresAt.Unix()
	return c
}

func (c *ClaimsBuilder) Build() *Claims {
	return c.claims
}

func NewClaimsBuilder() *ClaimsBuilder {
	a := new(ClaimsBuilder)
	a.claims = new(Claims)
	a.claims.StandardClaims = jwt.StandardClaims{}
	return a
}
