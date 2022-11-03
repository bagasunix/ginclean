package endpoints

import (
	"context"

	"github.com/bagasunix/ginclean/server/domains"
	"github.com/bagasunix/ginclean/server/endpoints/requests"
	"github.com/gin-gonic/gin"
)

type RoleEndpoint interface {
	CreateRole() gin.HandlerFunc
	ListRole() gin.HandlerFunc
	ViewRole() gin.HandlerFunc
	UpdateRole() gin.HandlerFunc
	DeleteRole() gin.HandlerFunc
}

type roleHandler struct {
	service domains.Service
}

// CreateRole implements RoleEndpoint
func (r *roleHandler) CreateRole() gin.HandlerFunc {
	return func(g *gin.Context) {
		var req requests.CreateRole
		if err := g.Bind(&req); err != nil {
			return
		}
		r.service.CreateRole(g, &req)
	}
}

// DeleteRole implements RoleEndpoint
func (r *roleHandler) DeleteRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requests.EntityId
		r.service.DeleteRole(ctx, &req)
	}
}

// ListRole implements RoleEndpoint
func (r *roleHandler) ListRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requests.BaseList
		r.service.ListRole(ctx, &req)
	}
}

// UpdateRole implements RoleEndpoint
func (r *roleHandler) UpdateRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requests.UpdateRole
		r.service.UpdateRole(context.Background(), &req)
	}
}

// ViewRole implements RoleEndpoint
func (r *roleHandler) ViewRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requests.EntityId
		r.service.ViewRole(ctx, &req)
	}
}

func NewRoleEndpoint(svc domains.Service) RoleEndpoint {
	return &roleHandler{service: svc}
}
