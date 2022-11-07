package entities

const (
	USER_STATUS_DISABLED = "DISABLED"
	USER_STATUS_ENABLED  = "ENABLED"
)

type User struct {
	Entity
	Account
	Name          string `json:"name"`
	Msisdn        string `json:"msisdn"`
	Role          string `json:"role"`
	VillageId     int64  `json:"villafe_id"`
	SubDistrictId int64  `json:"sub_distict_id"`
	CityId        int64  `json:"city_id"`
	ProvinceId    int64  `json:"province_id"`
	PostalCode    int64  `json:"postal_code"`
	RemarkAddress string `json:"remark_address"`
}

// UserBuilder Builder Object for User
type UserBuilder struct {
	EntityBuilder
	AccountBuilder
	name          string
	msisdn        string
	role          string
	villageId     int64
	subDistrictId int64
	cityId        int64
	provinceId    int64
	postalCode    int64
	remarkAddress string
}

// NewUserBuilder Constructor for UserBuilder
func NewUserBuilder() *UserBuilder {
	o := new(UserBuilder)
	return o
}

// Build Method which creates User
func (u *UserBuilder) Build() *User {
	o := new(User)
	o.Entity = *u.EntityBuilder.Build()
	o.Account = *u.AccountBuilder.Build()
	o.Name = u.name
	o.Msisdn = u.msisdn
	o.Role = u.role
	o.VillageId = u.villageId
	o.SubDistrictId = u.subDistrictId
	o.CityId = u.cityId
	o.ProvinceId = u.provinceId
	o.PostalCode = u.postalCode
	o.RemarkAddress = u.remarkAddress
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

// Setter method for the field role of type string in the object UserBuilder
func (u *UserBuilder) SetRole(role string) {
	u.role = role
}
