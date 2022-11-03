package usecases

import (
	"context"

	"github.com/bagasunix/ginclean/endpoints/requests"
	"github.com/bagasunix/ginclean/endpoints/responses"
)

type RoleService interface {
	CreateRole(ctx context.Context, req *requests.CreateRole) (res *responses.EntityId, err error)
	ListRole(ctx context.Context, req *requests.BaseList) (res *responses.ListEntity[entities.Role], err error)
	ViewRole(ctx context.Context, req *requests.EntityId) (res *responses.ViewEntity[*entities.Role], err error)
	UpdateRole(ctx context.Context, req *requests.UpdateRole) (res *responses.Empty, err error)
	DeleteRole(ctx context.Context, req *requests.EntityId) (res *responses.Empty, err error)
}
