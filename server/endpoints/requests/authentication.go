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
	Token string
}

func (s *Token) Validate() error {
	if validation.IsEmpty(s.Token) {
		return errors.ErrInvalidAttributes("token invalid")
	}
	return nil
}

func (s *Token) ToJSON() []byte {
	j, err := json.Marshal(s)
	errors.HandlerReturnedVoid(err)
	return j
}
