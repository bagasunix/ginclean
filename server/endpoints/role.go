package endpoints

import (
	"fmt"
	"net/http"

	"github.com/bagasunix/ginclean/pkg/errors"
	"github.com/bagasunix/ginclean/server/domains"
	"github.com/bagasunix/ginclean/server/endpoints/requests"
	"github.com/bagasunix/ginclean/server/endpoints/utils"
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
		dataRole, err := r.service.CreateRole(g, &req)
		if err != nil {
			g.JSON(http.StatusConflict, gin.H{"err": err.Error()})
			return
		}
		g.JSON(http.StatusCreated, dataRole)
	}
}

// DeleteRole implements RoleEndpoint
func (r *roleHandler) DeleteRole() gin.HandlerFunc {
	return func(g *gin.Context) {
		req, err := decodeByEntityIdEndpoint(g)
		if err != nil {
			g.JSON(http.StatusBadRequest, errors.NewBadRequest(err))
			return
		}
		dataRole, err := r.service.DeleteRole(g, req.(*requests.EntityId))
		if err != nil {
			g.JSON(http.StatusNotFound, errors.NewNotFound(fmt.Sprintf("%v", req.(*requests.EntityId).Id), err))
			return
		}
		g.JSON(http.StatusNoContent, dataRole)
	}
}

// ListRole implements RoleEndpointd
func (r *roleHandler) ListRole() gin.HandlerFunc {
	return func(g *gin.Context) {
		req, err := decodeBaseListEndpoint(g)
		if err != nil {
			g.JSON(http.StatusBadRequest, errors.NewBadRequest(err))
			return
		}
		dataRole, err := r.service.ListRole(g, req.(*requests.BaseList))
		if err != nil {
			return
		}
		g.JSON(http.StatusOK, dataRole)
	}
}

// UpdateRole implements RoleEndpoint
func (r *roleHandler) UpdateRole() gin.HandlerFunc {
	return func(g *gin.Context) {
		req, err := utils.DecodeByUpdateRoleEndpoint(g)
		if err != nil {
			g.JSON(http.StatusBadRequest, errors.NewBadRequest(err))
			return
		}
		// var dataRole *requests.Empty
		if _, err := r.service.UpdateRole(g, req.(*requests.UpdateRole)); err != nil {
			g.JSON(http.StatusConflict, errors.NewConflict(fmt.Sprint(req.(*requests.UpdateRole).Name), err))
			return
		}
		g.JSON(http.StatusOK, gin.H{})
	}
}

// ViewRole implements RoleEndpoint
func (r *roleHandler) ViewRole() gin.HandlerFunc {
	return func(g *gin.Context) {
		req, err := decodeByEntityIdEndpoint(g)
		if err != nil {
			g.JSON(http.StatusBadRequest, errors.NewBadRequest(err))
			return
		}
		dataRole, err := r.service.ViewRole(g, req.(*requests.EntityId))
		if err != nil {
			g.JSON(http.StatusNotFound, errors.NewNotFound(fmt.Sprintf("%v", req.(*requests.EntityId)), err))
			return
		}
		g.JSON(http.StatusOK, dataRole)
	}
}

func NewRoleEndpoint(svc domains.Service) RoleEndpoint {
	return &roleHandler{service: svc}
}
