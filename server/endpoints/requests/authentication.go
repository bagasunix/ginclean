package requests

import (
	"encoding/json"

	"github.com/bagasunix/ginclean/pkg/errors"
	validation "github.com/go-ozzo/ozzo-validation"
)

type SignInWithEmailPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *SignInWithEmailPassword) Validate() error {
	if validation.IsEmpty(s.Email) || validation.IsEmpty(s.Password) {
		return errors.ErrInvalidAttributes("email or password")
	}
	return nil
}

func (s *SignInWithEmailPassword) ToJSON() []byte {
	j, err := json.Marshal(s)
	errors.HandlerReturnedVoid(err)
	return j
}

type Token struct {
	Token string `json:"token"`
}

func (s *Token) Validate() error {
	if validation.IsEmpty(s.Token) {
		return errors.ErrInvalidAttributes("token")
	}
	return nil
}

func (s *Token) ToJSON() []byte {
	j, err := json.Marshal(s)
	errors.HandlerReturnedVoid(err)
	return j
}

// Builder Object for Token
type TokenBuilder struct {
	token string
}

// Constructor for TokenBuilder
func NewTokenBuilder() *TokenBuilder {
	o := new(TokenBuilder)
	return o
}

// Build Method which creates Token
func (b *TokenBuilder) Build() *Token {
	o := new(Token)
	o.Token = b.token
	return o
}

// Setter method for the field token of type string in the object TokenBuilder
func (t *TokenBuilder) SetToken(token string) {
	t.token = token
}
