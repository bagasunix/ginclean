package account

import (
	"context"

	"github.com/bagasunix/ginclean/pkg/errors"
	"github.com/bagasunix/ginclean/server/domains/data/models"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type gormProvider struct {
	db   *gorm.DB
	logs zap.Logger
}

// UpdateStatus implements AccountRepository
func (g *gormProvider) UpdateStatus(ctx context.Context, id uuid.UUID, stat bool) error {
	return errors.ErrSomethingWrong(g.logs, g.db.WithContext(ctx).Where("id = ?", id.String()).Update("isActive = ?", stat).Error)
}

// Delete implements AccountRepository
func (g *gormProvider) Delete(ctx context.Context, id uuid.UUID) error {
	return errors.ErrSomethingWrong(g.logs, g.db.WithContext(ctx).Delete(models.NewAccountBuilder().Build(), "id = ?", id.String()).Error)
}

// GetAll implements AccountRepository
func (g *gormProvider) GetAll(ctx context.Context, limit int64) (result models.SliceResult[models.Account]) {
	result.Error = errors.ErrRecordNotFound(g.logs, g.GetModelName(), g.db.WithContext(ctx).Limit(int(limit)).Find(&result.Value).Error)
	return result
}

// GetByEmail implements AccountRepository
func (g *gormProvider) GetByEmail(ctx context.Context, email string) (result models.SingleResult[*models.Account]) {
	result.Error = errors.ErrRecordNotFound(g.logs, g.GetModelName(), g.db.WithContext(ctx).First(&result.Value, email).Error)
	return result
}

// GetById implements AccountRepository
func (g *gormProvider) GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.Account]) {
	result.Error = errors.ErrRecordNotFound(g.logs, g.GetModelName(), g.db.WithContext(ctx).First(&result.Value, id).Error)
	return result
}

// Create implements AccountRepository
func (g *gormProvider) Create(ctx context.Context, user *models.Account) error {
	return errors.ErrDuplicateValue(g.logs, g.GetModelName(), g.db.WithContext(ctx).Create(user).Error)
}

// GetConnection implements AccountRepository
func (g *gormProvider) GetConnection() (T any) {
	return g.db
}

// GetModelName implements AccountRepository
func (g *gormProvider) GetModelName() string {
	return "Account"
}

func NewGorm(logs zap.Logger, db *gorm.DB) Repository {
	g := new(gormProvider)
	g.logs = logs
	g.db = db
	return g
}
