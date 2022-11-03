package inits

import (
	"github.com/bagasunix/ginclean/domains/data/models"
	"github.com/bagasunix/ginclean/pkg/errors"
	"gorm.io/gorm"
)

func GetTables() (tables []interface{}) {
	tables = append(tables, models.NewAccountBuilder().Build())
	tables = append(tables, models.NewRoleBuilder().Build())
	return tables
}
func Migrate(db *gorm.DB) {
	errors.HandlerWithOSExit(db.AutoMigrate(GetTables()...), "AutoMigrate")
}
