package requests

import (
	"encoding/json"

	"github.com/bagasunix/ginclean/pkg/errors"
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateUser struct {
	Name          string `json:"name"`
	Msisdn        string `json:"msisdn"`
	VillageId     int64  `json:"village_id"`
	SubDistrictId int64  `json:"sub_district_id"`
	CityId        int64  `json:"city_id"`
	ProvinceId    int64  `json:"province_id"`
	PostalCode    int64  `json:"postal_code"`
	RemarkAddress string `json:"remark_address"`
}

func (s *CreateUser) Validate() error {
	if validation.IsEmpty(s.Name) {
		return errors.ErrInvalidAttributes("name")
	}
	if validation.IsEmpty(s.Msisdn) {
		return errors.ErrInvalidAttributes("msisdn")
	}
	if validation.IsEmpty(s.VillageId) {
		return errors.ErrInvalidAttributes("village_id")
	}
	if validation.IsEmpty(s.SubDistrictId) {
		return errors.ErrInvalidAttributes("sub_district_id")
	}
	if validation.IsEmpty(s.CityId) {
		return errors.ErrInvalidAttributes("city_id")
	}
	if validation.IsEmpty(s.ProvinceId) {
		return errors.ErrInvalidAttributes("province_id")
	}
	if validation.IsEmpty(s.PostalCode) {
		return errors.ErrInvalidAttributes("postal_code")
	}
	if validation.IsEmpty(s.RemarkAddress) {
		return errors.ErrInvalidAttributes("remark_address")
	}
	return nil
}

func (s *CreateUser) ToJSON() []byte {
	j, err := json.Marshal(s)
	errors.HandlerReturnedVoid(err)
	return j
}
