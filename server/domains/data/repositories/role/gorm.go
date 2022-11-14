package role

import (
	"context"
	"fmt"

	"github.com/bagasunix/ginclean/pkg/errors"
	"github.com/bagasunix/ginclean/server/domains/data/models"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type gormProvider struct {
	logs *zap.Logger
	db   *gorm.DB
}

// GetByKeywords implements Repository
func (g *gormProvider) GetByKeywords(ctx context.Context, keywords string, limit int64) (result models.SliceResult[models.Role]) {
	a := fmt.Sprint('%', keywords, '%')
	result.Error = errors.ErrRecordNotFound(g.logs, g.GetModelName(), g.db.WithContext(ctx).Where("name like ?", a).Limit(int(limit)).Find(&result.Value).Error)
	return result
}

// GetAll implements Repository
func (g *gormProvider) GetAll(ctx context.Context, limit int64) (result models.SliceResult[models.Role]) {
	result.Error = errors.ErrRecordNotFound(g.logs, g.GetModelName(), g.db.WithContext(ctx).Limit(int(limit)).Find(&result.Value).Error)
	return result
}

// Delete implements Repository
func (g *gormProvider) Delete(ctx context.Context, id uuid.UUID) error {
	return errors.ErrSomethingWrong(g.logs, g.db.WithContext(ctx).Delete(models.NewRoleBuilder().Build(), "id = ?", id.String()).Error)
}

// Update implements Repository
func (g *gormProvider) Update(ctx context.Context, model *models.Role) error {
	return errors.ErrDuplicateValue(g.logs, g.GetModelName(), g.db.WithContext(ctx).Updates(model).Error)
}

// GetById implements Repository
func (g *gormProvider) GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.Role]) {
	result.Error = errors.ErrRecordNotFound(g.logs, g.GetModelName(), g.db.WithContext(ctx).Where("id = ?", id).First(&result.Value).Error)
	return result
}

// GetConnection implements Repository
func (g *gormProvider) GetConnection() (T any) {
	return g.db
}

// GetModelName implements Repository
func (g *gormProvider) GetModelName() string {
	return "Role"
}

// Create implements Repository
func (g *gormProvider) Create(ctx context.Context, role *models.Role) error {
	return errors.ErrDuplicateValue(g.logs, g.GetModelName(), g.db.WithContext(ctx).Create(role).Error)
}

func NewGorm(logs *zap.Logger, db *gorm.DB) Repository {
	g := new(gormProvider)
	g.db = db
	g.logs = logs
	return g
}
