package role

import (
	"context"

	"github.com/bagasunix/ginclean/server/domains/data/models"
	"github.com/bagasunix/ginclean/server/domains/data/repositories/base"
	"github.com/gofrs/uuid"
)

type Command interface {
	Create(ctx context.Context, model *models.Role) error
	Update(ctx context.Context, model *models.Role) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Query interface {
	GetAll(ctx context.Context, limit int64) (result models.SliceResult[models.Role])
	GetById(ctx context.Context, id uuid.UUID) (res models.SingleResult[*models.Role])
	GetByKeywords(ctx context.Context, keywords string, limit int64) (result models.SliceResult[models.Role])
}

type Repository interface {
	base.Repository
	Command
	Query
}
