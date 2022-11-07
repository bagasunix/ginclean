package requests

import (
	"encoding/json"

	"github.com/bagasunix/ginclean/pkg/errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofrs/uuid"
)

type CreateAccount struct {
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Role     uuid.UUID `json:"role"`
}

func (s *CreateAccount) Validate() error {
	if validation.IsEmpty(s.Email) {
		return errors.ErrInvalidAttributes("email")
	}
	if validation.IsEmpty(s.Password) {
		return errors.ErrInvalidAttributes("password")
	}
	if validation.IsEmpty(s.Role) {
		return errors.ErrInvalidAttributes("role")
	}
	return nil
}

func (s *CreateAccount) ToJSON() []byte {
	j, err := json.Marshal(s)
	errors.HandlerReturnedVoid(err)
	return j
}
