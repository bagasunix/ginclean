package account

import (
	"context"

	"github.com/bagasunix/ginclean.git/domains/data/models"
	"gorm.io/gorm"
)

type gormProvider struct {
	db *gorm.DB
}

// Create implements AccountRepository
func (g *gormProvider) Create(ctx context.Context, user *models.Account) error {
	panic("unimplemented")
}

// GetConnection implements AccountRepository
func (g *gormProvider) GetConnection() (T any) {
	return g.db
}

// GetModelName implements AccountRepository
func (g *gormProvider) GetModelName() string {
	return "Account"
}

func NewGorm(db *gorm.DB) AccountRepository {
	g := new(gormProvider)
	g.db = db
	return g
}
