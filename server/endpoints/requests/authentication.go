package requests

import (
	"encoding/json"

	"github.com/bagasunix/ginclean/pkg/errors"
	validation "github.com/go-ozzo/ozzo-validation"
)

type SignInWithUserNamePassword struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (s *SignInWithUserNamePassword) Validate() error {
	if validation.IsEmpty(s.UserName) || validation.IsEmpty(s.Password) {
		return errors.ErrInvalidAttributes("username or password")
	}
	return nil
}

func (s *SignInWithUserNamePassword) ToJSON() []byte {
	j, err := json.Marshal(s)
	errors.HandlerReturnedVoid(err)
	return j
}

type Token struct {
	Token string
}
