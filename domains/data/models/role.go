package models

type Role struct {
	BaseModel
	Name string `gorm:"size:100;uniqueIndex:idx_role_unique,sort:asc"`
}

// Builder Object for Role
type RoleBuilder struct {
	BaseModelBuilder
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
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.Name = b.name
	return o
}

// Setter method for the field name of type string in the object RoleBuilder
func (r *RoleBuilder) SetName(name string) {
	r.name = name
}
