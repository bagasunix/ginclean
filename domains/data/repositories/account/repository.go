package account

import (
	"context"

	"github.com/bagasunix/ginclean.git/domains/data/models"
	"github.com/bagasunix/ginclean.git/domains/data/repositories/base"
)

type Command interface {
	Create(ctx context.Context, user *models.Account) error
}

type Query interface {
}

type AccountRepository interface {
	Command
	Query
	base.Repository
}
