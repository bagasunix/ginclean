package inits

import (
	"context"
	"strconv"
	"time"

	"github.com/bagasunix/ginclean/pkg/db"
	envs "github.com/bagasunix/ginclean/pkg/env"
	"gorm.io/gorm"
)

func InitDb(ctx context.Context, configs *envs.Configs) *gorm.DB {

	MaxIdleConns, _ := strconv.Atoi(configs.MaxIdleConns)
	MaxOpenConns, _ := strconv.Atoi(configs.MaxOpenConns)
	ConnMaxLifetime, _ := time.ParseDuration(configs.ConnMaxLifetime + "m")

	configBuilder := db.NewDbConfigBuilder()
	configBuilder.SetDriver(configs.DbDriver)
	configBuilder.SetHost(configs.DbHost)
	configBuilder.SetPort(configs.DbPort)
	configBuilder.SetUser(configs.DbUsername)
	configBuilder.SetDatabaseName(configs.DbName)
	configBuilder.SetPassword(configs.DbPassword)
	configBuilder.SetSslMode(configs.SSLMode)
	configBuilder.SetTimeZone(configs.Timezone)
	configBuilder.SetMaxIdleConss(MaxIdleConns)
	configBuilder.SetMaxOpenConns(MaxOpenConns)
	configBuilder.SetConnMaxLifetime(ConnMaxLifetime)

	config := configBuilder.Build()
	return db.NewDB(ctx, config)
}
