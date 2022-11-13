package responses

import (
	"encoding/json"

	"github.com/bagasunix/ginclean/pkg/errors"
)

type SignIn struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

func (c *SignIn) ToJSON() []byte {
	j, err := json.Marshal(c)
	errors.HandlerReturnedVoid(err)
	return j
}

// Builder Object for SignIn
type SignInBuilder struct {
	token        string
	refreshToken string
}

// Constructor for SignInBuilder
func NewSignInBuilder() *SignInBuilder {
	o := new(SignInBuilder)
	return o
}

// Build Method which creates SignIn
func (b *SignInBuilder) Build() *SignIn {
	o := new(SignIn)
	o.Token = b.token
	o.RefreshToken = b.refreshToken
	return o
}

// Setter method for the field token of type string in the object SignInBuilder
func (s *SignInBuilder) SetToken(token string) *SignInBuilder {
	s.token = token
	return s
}

// Setter method for the field refreshToken of type string in the object SignInBuilder
func (s *SignInBuilder) SetRefreshToken(refreshToken string) *SignInBuilder {
	s.refreshToken = refreshToken
	return s
}

type RefreshToken struct {
	Token string `json:"token"`
}

func (c *RefreshToken) ToJSON() []byte {
	j, err := json.Marshal(c)
	errors.HandlerReturnedVoid(err)
	return j
}

// Builder Object for RefreshToken
type RefreshTokenBuilder struct {
	token string
}

// Constructor for RefreshTokenBuilder
func NewRefreshTokenBuilder() *RefreshTokenBuilder {
	o := new(RefreshTokenBuilder)
	return o
}

// Build Method which creates RefreshToken
func (b *RefreshTokenBuilder) Build() *RefreshToken {
	o := new(RefreshToken)
	o.Token = b.token
	return o
}

// Setter method for the field token of type string in the object RefreshTokenBuilder
func (r *RefreshTokenBuilder) SetToken(token string) {
	r.token = token
}
