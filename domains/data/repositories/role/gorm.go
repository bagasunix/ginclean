package role

import (
	"context"
	"fmt"

	"github.com/bagasunix/ginclean/domains/data/models"
	"github.com/bagasunix/ginclean/pkg/errors"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type gormProvider struct {
	db *gorm.DB
}

// GetByKeywords implements Repository
func (g *gormProvider) GetByKeywords(ctx context.Context, keywords string, limit int64) (result models.SliceResult[models.Role]) {
	a := fmt.Sprint('%', keywords, '%')
	result.Error = errors.NewNotFound(g.GetModelName(), g.db.WithContext(ctx).Where("name like ?", a).Limit(int(limit)).Find(&result.Value).Error.Error())
	return result
}

// GetAll implements Repository
func (g *gormProvider) GetAll(ctx context.Context, limit int64) (result models.SliceResult[models.Role]) {
	result.Error = errors.NewNotFound(g.GetModelName(), g.db.WithContext(ctx).Limit(int(limit)).Find(&result.Value).Error.Error())
	return result
}

// Delete implements Repository
func (g *gormProvider) Delete(ctx context.Context, id uuid.UUID) error {
	return errors.NewBadRequest(g.db.WithContext(ctx).Delete(models.NewRoleBuilder().Build(), "id = ?", id.String()).Error.Error())
}

// Update implements Repository
func (g *gormProvider) Update(ctx context.Context, model *models.Role) error {
	return errors.NewConflict(g.GetModelName(), g.db.WithContext(ctx).Updates(model).Error.Error())
}

// GetById implements Repository
func (g *gormProvider) GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.Role]) {
	result.Error = errors.NewNotFound(g.GetModelName(), g.db.WithContext(ctx).Where("id = ?", id).First(&result.Value).Error.Error())
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
	return errors.NewConflict(g.GetModelName(), g.db.WithContext(ctx).Create(role).Error.Error())
}

func NewGorm(db *gorm.DB) Repository {
	g := new(gormProvider)
	g.db = db
	return g
}
