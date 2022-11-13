package jwt

import (
	"time"

	"github.com/bagasunix/ginclean/server/domains/entities"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	User   *entities.Account `json:"account,omitempty"`
	Client *entities.Client  `json:"client,omitempty"`
	jwt.StandardClaims
}

type ClaimsBuilder struct {
	claims *Claims
}

func (c *ClaimsBuilder) Client(client *entities.Client) *ClaimsBuilder {
	c.claims.Client = client
	return c
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
