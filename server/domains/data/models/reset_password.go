package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type ResetPassword struct {
	BaseModel
	AccountId    uuid.UUID
	Account      *Account `gorm:"foreignKey:AccountId;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	ChangeAt     time.Time
	ResetToken   string
	ResetExpires string
}

// Builder Object for ResetPassword
type ResetPasswordBuilder struct {
	BaseModelBuilder
	accountId    uuid.UUID
	account      *Account
	changeAt     time.Time
	resetToken   string
	resetExpires string
}

// Constructor for ResetPasswordBuilder
func NewResetPasswordBuilder() *ResetPasswordBuilder {
	o := new(ResetPasswordBuilder)
	return o
}

// Build Method which creates ResetPassword
func (b *ResetPasswordBuilder) Build() *ResetPassword {
	o := new(ResetPassword)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.AccountId = b.accountId
	o.Account = b.account
	o.ChangeAt = b.changeAt
	o.ResetToken = b.resetToken
	o.ResetExpires = b.resetExpires
	return o
}

// Setter method for the field accountId of type uuid.UUID in the object ResetPasswordBuilder
func (r *ResetPasswordBuilder) SetAccountId(accountId uuid.UUID) {
	r.accountId = accountId
}

// Setter method for the field account of type *Account in the object ResetPasswordBuilder
func (r *ResetPasswordBuilder) SetAccount(account *Account) {
	r.account = account
}

// Setter method for the field changeAt of type time.Time in the object ResetPasswordBuilder
func (r *ResetPasswordBuilder) SetChangeAt(changeAt time.Time) {
	r.changeAt = changeAt
}

// Setter method for the field resetToken of type string in the object ResetPasswordBuilder
func (r *ResetPasswordBuilder) SetResetToken(resetToken string) {
	r.resetToken = resetToken
}

// Setter method for the field resetExpires of type string in the object ResetPasswordBuilder
func (r *ResetPasswordBuilder) SetResetExpires(resetExpires string) {
	r.resetExpires = resetExpires
}
