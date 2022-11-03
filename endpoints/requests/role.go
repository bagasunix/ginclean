package requests

import (
	"encoding/json"

	"github.com/bagasunix/ginclean/pkg/errors"
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateRole struct {
	Name string `json:"name"`
}

func (s *CreateRole) Validate() error {
	if validation.IsEmpty(s.Name) {
		return errors.ErrInvalidAttributes("name")
	}
	return nil
}

func (s *CreateRole) ToJSON() []byte {
	j, err := json.Marshal(s)
	errors.HandlerReturnedVoid(err)
	return j
}

type UpdateRole struct {
	EntityId
	Name string `json:"name"`
}

func (s *UpdateRole) Validate() error {
	if validation.IsEmpty(s.Name) {
		return errors.ErrInvalidAttributes("name")
	}
	return nil
}

func (s *UpdateRole) ToJSON() []byte {
	j, err := json.Marshal(s)
	errors.HandlerReturnedVoid(err)
	return j
}
