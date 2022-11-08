package responses

import (
	"encoding/json"

	"github.com/bagasunix/ginclean/pkg/errors"
)

type Role struct {
	EntityId
	Name string `json:"name"`
}

func (c *Role) ToJSON() []byte {
	j, err := json.Marshal(c)
	errors.HandlerReturnedVoid(err)
	return j
}

// Builder Object for Role
type RoleBuilder struct {
	EntityIdBuilder
	name string
}

// Constructor for RoleBuilder
func NewRoleBuilder() *RoleBuilder {
	o := new(RoleBuilder)
	return o
}

// Build Method which creates Role
func (b *RoleBuilder) Build() *Role {
	o := new(Role)
	o.EntityId = *b.EntityIdBuilder.Build()
	o.Name = b.name
	return o
}

// Setter method for the field name of type string in the object RoleBuilder
func (r *RoleBuilder) SetName(name string) {
	r.name = name
}
