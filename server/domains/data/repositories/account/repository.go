package account

import (
	"context"

	"github.com/bagasunix/ginclean/server/domains/data/models"
	"github.com/bagasunix/ginclean/server/domains/data/repositories/base"
	"github.com/gofrs/uuid"
)

type Command interface {
	Create(ctx context.Context, user *models.Account) error
	Delete(ctx context.Context, id uuid.UUID) error
	UpdateStatus(ctx context.Context, model *models.Account) error
}

type Query interface {
	GetAll(ctx context.Context, limit int64) (result models.SliceResult[models.Account])
	GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.Account])
	GetByEmail(ctx context.Context, email string) (result models.SingleResult[*models.Account])
	GetByKeywords(ctx context.Context, keywords string, limit int64) (result models.SliceResult[models.Account])
}

type Repository interface {
	Command
	Query
	base.Repository
}
