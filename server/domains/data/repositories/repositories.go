package repositories

import (
	"github.com/bagasunix/ginclean/server/domains/data/repositories/account"
	"github.com/bagasunix/ginclean/server/domains/data/repositories/role"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repositories interface {
	GetRole() role.Repository
	GetAccount() account.Repository
}

type repo struct {
	role    role.Repository
	account account.Repository
}

// GetAccount implements Repositories
func (r *repo) GetAccount() account.Repository {
	return r.account
}

// GetRole implements Repositories
func (r *repo) GetRole() role.Repository {
	return r.role
}

func New(logs zap.Logger, db *gorm.DB) Repositories {
	rs := new(repo)
	rs.role = role.NewGorm(logs, db)
	rs.account = account.NewGorm(logs, db)
	return rs
}
