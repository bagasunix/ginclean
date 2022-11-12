package models

import validation "github.com/go-ozzo/ozzo-validation"

type Address struct {
	VillageId int64    `gorm:"index"`
	Village   *Village `gorm:"foreignKey:VillageId"`
	*Coordinate
	StreetAddress string `gorm:"size:300"`
	PostalCode    string `gorm:"size:8"`
}

// AddressBuilder Builder Object for Address
type AddressBuilder struct {
	CoordinateBuilder
	villageId     int64
	village       *Village
	streetAddress string
	postalCode    string
}

// NewAddressBuilder Constructor for AddressBuilder
func NewAddressBuilder() *AddressBuilder {
	o := new(AddressBuilder)
	return o
}

// Build Method which creates Address
func (b *AddressBuilder) Build() *Address {
	o := new(Address)
	o.StreetAddress = b.streetAddress
	o.PostalCode = b.postalCode
	o.VillageId = b.villageId
	o.Village = b.village
	o.Coordinate = b.CoordinateBuilder.Build()
	return o
}

// SetVillage Setter method for the field village of type *Village in the object AddressBuilder
func (b *AddressBuilder) SetVillage(village *Village) {
	if validation.IsEmpty(village) {
		return
	}
	if validation.IsEmpty(village.Id) {
		return
	}
	b.village = village
	b.villageId = village.Id
}

// SetStreetAddress Setter method for the field streetAddress of type string in the object AddressBuilder
func (b *AddressBuilder) SetStreetAddress(streetAddress string) {
	b.streetAddress = streetAddress
}

// SetPostalCode Setter method for the field postalCode of type string in the object AddressBuilder
func (b *AddressBuilder) SetPostalCode(postalCode string) {
	b.postalCode = postalCode
}
