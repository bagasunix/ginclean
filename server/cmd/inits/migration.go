package inits

import (
	"github.com/bagasunix/ginclean/pkg/errors"
	"github.com/bagasunix/ginclean/server/domains/data/models"
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
