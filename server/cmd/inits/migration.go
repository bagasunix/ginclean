package inits

import (
	"github.com/bagasunix/ginclean/pkg/errors"
	"github.com/bagasunix/ginclean/server/domains/data/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GetTables() (tables []interface{}) {
	tables = append(tables, models.NewAccountBuilder().Build())
	tables = append(tables, models.NewRoleBuilder().Build())
	tables = append(tables, models.NewUserBuilder().Build())
	tables = append(tables, models.NewRefershTokenBuilder().Build())
	tables = append(tables, models.NewCountryBuilder().Build())
	tables = append(tables, models.NewProvinceBuilder().Build())
	tables = append(tables, models.NewCityBuilder().Build())
	tables = append(tables, models.NewSubDistrictBuilder().Build())
	tables = append(tables, models.NewVillageBuilder().Build())
	return tables
}
func Migrate(logs zap.Logger, db *gorm.DB) {
	errors.HandlerWithOSExit(logs, db.AutoMigrate(GetTables()...), "AutoMigrate")
}
