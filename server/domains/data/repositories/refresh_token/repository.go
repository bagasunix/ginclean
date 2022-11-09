package refreshtoken

import (
	"context"

	"github.com/bagasunix/ginclean/server/domains/data/models"
	"github.com/bagasunix/ginclean/server/domains/data/repositories/base"
	"github.com/gofrs/uuid"
)

type Command interface {
	Create(ctx context.Context, m *models.RefershToken) error
	Delete(ctx context.Context, userId uuid.UUID) error
}

type Query interface {
}

type Repository interface {
	base.Repository
	Command
	Query
}
