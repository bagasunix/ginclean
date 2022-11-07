package entities

type Account struct {
	Entity
	Email    string `json:"name"`
	Role     *Role  `json:"role"`
	IsActive bool   `json:"is_active"`
}

// Builder Object for Account
type AccountBuilder struct {
	EntityBuilder
	email    string
	Role     Role
	isActive bool
}

// Constructor for AccountBuilder
func NewAccountBuilder() *AccountBuilder {
	o := new(AccountBuilder)
	return o
}

// Build Method which creates Account
func (b *AccountBuilder) Build() *Account {
	o := new(Account)
	o.Entity = *b.EntityBuilder.Build()
	o.Role = &b.Role
	o.Email = b.email
	o.IsActive = b.isActive
	return o
}

// Setter method for the field email of type string in the object AccountBuilder
func (a *AccountBuilder) SetEmail(email string) {
	a.email = email
}

// Setter method for the field isActive of type bool in the object AccountBuilder
func (a *AccountBuilder) SetIsActive(isActive bool) {
	a.isActive = isActive
}
