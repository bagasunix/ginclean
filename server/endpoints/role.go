package endpoints

import (
	"net/http"

	"github.com/bagasunix/ginclean/pkg/errors"
	"github.com/bagasunix/ginclean/server/domains"
	"github.com/bagasunix/ginclean/server/endpoints/requests"
	"github.com/gin-gonic/gin"
)

const (
	CREATE_ROLE = "CreateRole"
	UPDATE_ROLE = "UpdateRole"
	DELETE_ROLE = "DeleteRole"
	LIST_ROLE   = "ListRole"
	VIEW_ROLE   = "ViewRole"
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
			g.JSON(http.StatusBadRequest, errors.NewBadRequest(err))
			return
		}
		dataRole, err := r.service.CreateRole(g.Request.Context(), &req)
		if err != nil {
			return
		}
		g.JSON(http.StatusOK, dataRole)
		return
	}
}

// DeleteRole implements RoleEndpoint
func (r *roleHandler) DeleteRole() gin.HandlerFunc {
	return func(g *gin.Context) {
		var req requests.EntityId
		if err := g.Bind(&req); err != nil {
			g.JSON(http.StatusBadRequest, errors.NewBadRequest(err))
			return
		}
		dataRole, err := r.service.DeleteRole(g, &req)
		if err != nil {
			return
		}
		g.JSON(http.StatusOK, dataRole)
		return
	}
}

// ListRole implements RoleEndpointd
func (r *roleHandler) ListRole() gin.HandlerFunc {
	return func(g *gin.Context) {
		var req requests.BaseList
		if err := g.Bind(&req); err != nil {
			g.JSON(http.StatusBadRequest, errors.NewBadRequest(err))
			return
		}
		dataRole, err := r.service.ListRole(g, &req)
		if err != nil {
			return
		}
		g.JSON(http.StatusOK, dataRole)
		return
	}
}

// UpdateRole implements RoleEndpoint
func (r *roleHandler) UpdateRole() gin.HandlerFunc {
	return func(g *gin.Context) {
		var req requests.UpdateRole
		if err := g.Bind(&req); err != nil {
			g.JSON(http.StatusBadRequest, errors.NewBadRequest(err))
			return
		}
		dataRole, err := r.service.UpdateRole(g, &req)
		if err != nil {
			return
		}
		g.JSON(http.StatusOK, dataRole)
		return
	}
}

// ViewRole implements RoleEndpoint
func (r *roleHandler) ViewRole() gin.HandlerFunc {
	return func(g *gin.Context) {
		var req requests.EntityId
		if err := g.Bind(&req); err != nil {
			g.JSON(http.StatusBadRequest, errors.NewBadRequest(err))
			return
		}
		dataRole, err := r.service.ViewRole(g, &req)
		if err != nil {
			return
		}
		g.JSON(http.StatusOK, dataRole)
		return
	}
}

func NewRoleEndpoint(svc domains.Service) RoleEndpoint {
	return &roleHandler{service: svc}
}
