package db

import (
	"context"
	"time"

	"github.com/bagasunix/ginclean/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func NewDB(ctx context.Context, dbConfig *DbConfig) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbConfig.GetDSN()), &gorm.Config{SkipDefaultTransaction: true, PrepareStmt: true})
	errors.HandlerWithOSExit(err, "init", "database", "config", dbConfig.GetDSN())
	errors.HandlerWithOSExit(db.WithContext(ctx).Use(dbresolver.Register(dbresolver.Config{}).SetMaxIdleConns(dbConfig.MaxIdleConns).SetMaxOpenConns(dbConfig.MaxOpenConns).SetConnMaxLifetime(5*time.Minute)), "db_resolver")
	return db
}
