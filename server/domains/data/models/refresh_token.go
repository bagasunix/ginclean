package models

import "github.com/gofrs/uuid"

type RefershToken struct {
	BaseModel
	UserId uuid.UUID
	User   *Account `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	Token  string
}

// Builder Object for RefershToken
type RefershTokenBuilder struct {
	BaseModelBuilder
	userId uuid.UUID
	user   *Account
	token  string
}

// Constructor for RefershTokenBuilder
func NewRefershTokenBuilder() *RefershTokenBuilder {
	o := new(RefershTokenBuilder)
	return o
}

// Build Method which creates RefershToken
func (b *RefershTokenBuilder) Build() *RefershToken {
	o := new(RefershToken)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.UserId = b.userId
	o.User = b.user
	o.Token = b.token
	return o
}

// Setter method for the field userId of type uuid.UUID in the object RefershTokenBuilder
func (r *RefershTokenBuilder) SetUserId(userId uuid.UUID) {
	r.userId = userId
}

// Setter method for the field user of type *Account in the object RefershTokenBuilder
func (r *RefershTokenBuilder) SetUser(user *Account) {
	r.user = user
}

// Setter method for the field token of type string in the object RefershTokenBuilder
func (r *RefershTokenBuilder) SetToken(token string) {
	r.token = token
}
