package repositories

import (
	"github.com/bagasunix/ginclean/server/domains/data/repositories/account"
	refreshtoken "github.com/bagasunix/ginclean/server/domains/data/repositories/refresh_token"
	"github.com/bagasunix/ginclean/server/domains/data/repositories/role"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repositories interface {
	GetRole() role.Repository
	GetAccount() account.Repository
	GetRefreshToken() refreshtoken.Repository
}

type repo struct {
	role         role.Repository
	account      account.Repository
	refreshToken refreshtoken.Repository
}

// GetRefreshToken implements Repositories
func (r *repo) GetRefreshToken() refreshtoken.Repository {
	return r.refreshToken
}

// GetAccount implements Repositories
func (r *repo) GetAccount() account.Repository {
	return r.account
}

// GetRole implements Repositories
func (r *repo) GetRole() role.Repository {
	return r.role
}

func New(logs *zap.Logger, db *gorm.DB) Repositories {
	rs := new(repo)
	rs.role = role.NewGorm(logs, db)
	rs.account = account.NewGorm(logs, db)
	rs.refreshToken = refreshtoken.NewGorm(logs, db)
	return rs
}
