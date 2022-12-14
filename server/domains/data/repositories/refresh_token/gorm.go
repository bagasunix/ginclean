package refreshtoken

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
	logs *zap.Logger
}

// Delete implements Repository
func (g *gormProvider) Delete(ctx context.Context, userId uuid.UUID) error {
	return errors.ErrSomethingWrong(g.logs, g.db.WithContext(ctx).Delete(models.NewRefershTokenBuilder().Build(), "user_id = ?", userId).Error)
}

// GetConnection implements Repository
func (g *gormProvider) GetConnection() (T any) {
	return g.db
}

// GetModelName implements Repository
func (g *gormProvider) GetModelName() string {
	return "RefreshToken"
}

// CreateRefershToken implements Repository
func (g *gormProvider) Create(ctx context.Context, m *models.RefershToken) error {
	return errors.ErrSomethingWrong(g.logs, g.db.WithContext(ctx).Create(m).Error)
}

func NewGorm(logs *zap.Logger, db *gorm.DB) Repository {
	g := new(gormProvider)
	g.logs = logs
	g.db = db
	return g
}
