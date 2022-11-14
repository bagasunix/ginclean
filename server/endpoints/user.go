package endpoints

import (
	"net/http"

	"github.com/bagasunix/ginclean/pkg/errors"
	"github.com/bagasunix/ginclean/server/domains"
	"github.com/bagasunix/ginclean/server/endpoints/requests"
	"github.com/bagasunix/ginclean/server/endpoints/utils"
	"github.com/gin-gonic/gin"
)

type UserEndpoint interface {
	CreateUser() gin.HandlerFunc
	ListAccount() gin.HandlerFunc
	DeleteAccount() gin.HandlerFunc
	DisableAccount() gin.HandlerFunc
	ViewAccount() gin.HandlerFunc
}

type userHandler struct {
	service domains.Service
}

// ViewAccount implements UserEndpoint
func (u *userHandler) ViewAccount() gin.HandlerFunc {
	return func(g *gin.Context) {
		req, err := decodeByEntityIdEndpoint(g)
		if err != nil {
			utils.EncodeError(g, err, g.Writer)
			return
		}
		dataRole, err := u.service.ViewAccountByID(g, req.(*requests.EntityId))
		if err != nil {
			g.JSON(http.StatusBadRequest, errors.NewBadRequest(err))
			return
		}
		g.JSON(http.StatusOK, dataRole)
	}
}

// DisableAccount implements UserEndpoint
func (u *userHandler) DisableAccount() gin.HandlerFunc {
	return func(g *gin.Context) {
		var req requests.DisableAccount
		if err := g.Bind(&req); err != nil {
			utils.EncodeError(g, err, g.Writer)
			return
		}
		dataAccount, err := u.service.DisableAccount(g, &req)
		if err != nil {
			utils.EncodeError(g, err, g.Writer)
			return
		}
		g.JSON(http.StatusOK, dataAccount)
	}
}

// DeleteAccount implements UserEndpoint
func (u *userHandler) DeleteAccount() gin.HandlerFunc {
	return func(g *gin.Context) {
		req, err := decodeByEntityIdEndpoint(g)
		if err != nil {
			utils.EncodeError(g, err, g.Writer)
			return
		}
		dataAccount, err := u.service.DeleteAccount(g, req.(*requests.EntityId))
		if err != nil {
			utils.EncodeError(g, err, g.Writer)
			return
		}
		g.JSON(http.StatusNoContent, dataAccount)
	}
}

// ListAccount implements UserEndpoint
func (u *userHandler) ListAccount() gin.HandlerFunc {
	return func(g *gin.Context) {
		req, err := decodeBaseListEndpoint(g)
		if err != nil {
			utils.EncodeError(g, err, g.Writer)
			return
		}
		dataAccount, err := u.service.ListAccount(g, req.(*requests.BaseList))
		if err != nil {
			utils.EncodeError(g, err, g.Writer)
			return
		}
		g.JSON(http.StatusOK, dataAccount)
	}
}

// CreateUser implements UserEndpoint
func (u *userHandler) CreateUser() gin.HandlerFunc {
	return func(g *gin.Context) {
		var req requests.CreateAccount
		if err := g.Bind(&req); err != nil {
			utils.EncodeError(g, err, g.Writer)
			return
		}
		dataAccount, err := u.service.CreateAccount(g, &req)
		if err != nil {
			utils.EncodeError(g, err, g.Writer)
			return
		}
		g.JSON(http.StatusCreated, dataAccount)
	}
}

func NewUserEndpoint(svc domains.Service) UserEndpoint {
	return &userHandler{service: svc}
}
