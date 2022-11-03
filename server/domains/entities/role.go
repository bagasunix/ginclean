package entities

type Role struct {
	Entity
	Name string `json:"name"`
}

// Builder Object for Role
type RoleBuilder struct {
	EntityBuilder
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
	o.Entity = *b.EntityBuilder.Build()
	o.Name = b.name
	return o
}

// Setter method for the field name of type string in the object RoleBuilder
func (r *RoleBuilder) SetName(name string) {
	r.name = name
}
