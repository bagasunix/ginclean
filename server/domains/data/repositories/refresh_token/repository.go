package refreshtoken

import (
	"context"

	"github.com/bagasunix/ginclean/server/domains/data/models"
	"github.com/bagasunix/ginclean/server/domains/data/repositories/base"
)

type Command interface {
	CreateRefershToken(ctx context.Context, m *models.RefershToken) error
}

type Query interface {
}

type Repository interface {
	base.Repository
	Command
	Query
}
