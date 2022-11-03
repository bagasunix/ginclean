package endpoints

import (
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
func (*roleHandler) DeleteRole() gin.HandlerFunc {
	panic("unimplemented")
}

// ListRole implements RoleEndpoint
func (*roleHandler) ListRole() gin.HandlerFunc {
	panic("unimplemented")
}

// UpdateRole implements RoleEndpoint
func (*roleHandler) UpdateRole() gin.HandlerFunc {
	panic("unimplemented")
}

// ViewRole implements RoleEndpoint
func (*roleHandler) ViewRole() gin.HandlerFunc {
	panic("unimplemented")
}

func NewRoleEndpoint(svc domains.Service) RoleEndpoint {
	return &roleHandler{service: svc}
}
