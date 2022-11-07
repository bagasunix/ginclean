package db

import (
	"context"
	"time"

	"github.com/bagasunix/ginclean/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func NewDB(ctx context.Context, logs zap.Logger, dbConfig *DbConfig) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbConfig.GetDSN()), &gorm.Config{SkipDefaultTransaction: true, PrepareStmt: true})
	errors.HandlerWithOSExit(logs, err, "init", "database", "config", dbConfig.GetDSN())
	errors.HandlerWithOSExit(logs, db.WithContext(ctx).Use(dbresolver.Register(dbresolver.Config{}).SetMaxIdleConns(dbConfig.MaxIdleConns).SetMaxOpenConns(dbConfig.MaxOpenConns).SetConnMaxLifetime(5*time.Minute)), "db_resolver")
	return db
}
