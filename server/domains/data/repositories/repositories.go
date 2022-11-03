package repositories

import (
	"github.com/bagasunix/ginclean/domains/data/repositories/role"
	"gorm.io/gorm"
)

type Repositories interface {
	GetRole() role.Repository
}

type repo struct {
	role role.Repository
}

// GetRole implements Repositories
func (r *repo) GetRole() role.Repository {
	return r.role
}

func New(db *gorm.DB) Repositories {
	rs := new(repo)
	rs.role = role.NewGorm(db)
	return rs
}
