package models

import "github.com/gofrs/uuid"

type Account struct {
	BaseModel
	Email    string `gorm:"size:100;uniqueIndex:idx_account_unique"`
	Password string
	RoleId   uuid.UUID
	Role     *Role `gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	IsActive bool
}

// Builder Object for User
type accountBuilder struct {
	BaseModelBuilder
	email    string
	password string
	roleId   uuid.UUID
	role     *Role
	isActive bool
}

// Constructor for accountBuilder
func NewAccountBuilder() *accountBuilder {
	o := new(accountBuilder)
	return o
}

// Build Method which creates User
func (b *accountBuilder) Build() *Account {
	o := new(Account)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.Email = b.email
	o.Password = b.password
	o.RoleId = b.roleId
	o.Role = b.role
	o.IsActive = b.isActive
	return o
}

// Setter method for the field email of type string in the object accountBuilder
func (u *accountBuilder) SetEmail(email string) {
	u.email = email
}

// Setter method for the field password of type string in the object accountBuilder
func (u *accountBuilder) SetPassword(password string) {
	u.password = password
}

// Setter method for the field roleId of type uuid.UUID in the object accountBuilder
func (u *accountBuilder) SetRoleId(roleId uuid.UUID) {
	u.roleId = roleId
}

// Setter method for the field role of type *Role in the object accountBuilder
func (u *accountBuilder) SetRole(role *Role) {
	u.role = role
}

// Setter method for the field isActive of type bool in the object accountBuilder
func (u *accountBuilder) SetIsActive(isActive bool) {
	u.isActive = isActive
}
