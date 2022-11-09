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
	tables = append(tables, models.NewRefershTokenBuilder().Build())
	return tables
}
func Migrate(logs zap.Logger, db *gorm.DB) {
	errors.HandlerWithOSExit(logs, db.AutoMigrate(GetTables()...), "AutoMigrate")
}
