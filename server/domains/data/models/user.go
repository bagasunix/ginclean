package models

import "github.com/gofrs/uuid"

type User struct {
	BaseModel
	AccountID     uuid.UUID
	Account       *Account `gorm:"foreignKey:AccountID;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	Name          string   `gorm:"size:100;index:,sort:asc,option:CONCURRENTLY;"`
	Msisdn        string   `gorm:"size:15;uniqueIndex:idx_msisdn_unique"`
	VillageId     int64
	Village       *Village `gorm:"foreignKey:VillageId;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	SubDistrictId int64
	SubDistrict   *SubDistrict `gorm:"foreignKey:SubDistrictId;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	CityId        int64
	City          *City `gorm:"foreignKey:CityId;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	ProvinceId    int64
	Province      *Province `gorm:"foreignKey:ProvinceId;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	CountryId     int64
	Country       *Country `gorm:"foreignKey:CountryId;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	PostalCode    int64
	RemarkAddress string
}

// Builder Object for User
type UserBuilder struct {
	BaseModelBuilder
	accountID     uuid.UUID
	account       *Account
	name          string
	msisdn        string
	villageId     int64
	village       *Village
	subDistrictId int64
	subDistrict   *SubDistrict
	cityId        int64
	city          *City
	provinceId    int64
	province      *Province
	countryId     int64
	country       *Country
	postalCode    int64
	remarkAddress string
}

// Constructor for UserBuilder
func NewUserBuilder() *UserBuilder {
	o := new(UserBuilder)
	return o
}

// Build Method which creates User
func (b *UserBuilder) Build() *User {
	o := new(User)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.AccountID = b.accountID
	o.Account = b.account
	o.Name = b.name
	o.Msisdn = b.msisdn
	o.VillageId = b.villageId
	o.Village = b.village
	o.SubDistrictId = b.subDistrictId
	o.SubDistrict = b.subDistrict
	o.CityId = b.cityId
	o.City = b.city
	o.ProvinceId = b.provinceId
	o.Province = b.province
	o.CountryId = b.countryId
	o.Country = b.country
	o.PostalCode = b.postalCode
	o.RemarkAddress = b.remarkAddress
	return o
}

// Setter method for the field name of type string in the object UserBuilder
func (u *UserBuilder) SetName(name string) {
	u.name = name
}

// Setter method for the field msisdn of type string in the object UserBuilder
func (u *UserBuilder) SetMsisdn(msisdn string) {
	u.msisdn = msisdn
}

// Setter method for the field villageId of type int64 in the object UserBuilder
func (u *UserBuilder) SetVillageId(villageId int64) {
	u.villageId = villageId
}

// Setter method for the field subDistrictId of type int64 in the object UserBuilder
func (u *UserBuilder) SetSubDistrictId(subDistrictId int64) {
	u.subDistrictId = subDistrictId
}

// Setter method for the field cityId of type int64 in the object UserBuilder
func (u *UserBuilder) SetCityId(cityId int64) {
	u.cityId = cityId
}

// Setter method for the field provinceId of type int64 in the object UserBuilder
func (u *UserBuilder) SetProvinceId(provinceId int64) {
	u.provinceId = provinceId
}

// Setter method for the field postalCode of type int64 in the object UserBuilder
func (u *UserBuilder) SetPostalCode(postalCode int64) {
	u.postalCode = postalCode
}

// Setter method for the field remarkAddress of type string in the object UserBuilder
func (u *UserBuilder) SetRemarkAddress(remarkAddress string) {
	u.remarkAddress = remarkAddress
}

// Setter method for the field account of type *Account in the object UserBuilder
func (u *UserBuilder) SetAccount(account *Account) {
	u.account = account
}

// Setter method for the field accountID of type uuid.UUID in the object UserBuilder
func (u *UserBuilder) SetAccountID(accountID uuid.UUID) {
	u.accountID = accountID
}

// Setter method for the field village of type *Village in the object UserBuilder
func (u *UserBuilder) SetVillage(village *Village) {
	u.village = village
}

// Setter method for the field subDistrict of type *SubDistrict in the object UserBuilder
func (u *UserBuilder) SetSubDistrict(subDistrict *SubDistrict) {
	u.subDistrict = subDistrict
}

// Setter method for the field city of type *City in the object UserBuilder
func (u *UserBuilder) SetCity(city *City) {
	u.city = city
}

// Setter method for the field province of type *Province in the object UserBuilder
func (u *UserBuilder) SetProvince(province *Province) {
	u.province = province
}


// Setter method for the field countryId of type int64 in the object UserBuilder
func (u *UserBuilder) SetCountryId(countryId int64) {		
	u.countryId = countryId
}

// Setter method for the field country of type *Country in the object UserBuilder
func (u *UserBuilder) SetCountry(country *Country) {		
	u.country = country
}