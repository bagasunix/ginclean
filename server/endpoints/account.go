package endpoints

import (
	"net/http"

	"github.com/bagasunix/ginclean/server/domains"
	"github.com/bagasunix/ginclean/server/endpoints/requests"
	"github.com/bagasunix/ginclean/server/endpoints/utils"
	"github.com/gin-gonic/gin"
)

type AccountEndpoint interface {
	CreateAccount() gin.HandlerFunc
	ListAccount() gin.HandlerFunc
	ViewAccount() gin.HandlerFunc
	UpdateAccount() gin.HandlerFunc
	UpdateMultiAccount() gin.HandlerFunc
	DeleteAccount() gin.HandlerFunc
}

type accountHandler struct {
	service domains.Service
}

// CreateAccount implements AccountEndpoint
func (a *accountHandler) CreateAccount() gin.HandlerFunc {
	return func(g *gin.Context) {
		var req requests.CreateAccount
		if err := g.Bind(&req); err != nil {
			utils.EncodeError(g, err, g.Writer)
			return
		}
		dataAccount, err := a.service.CreateAccount(g, &req)
		if err != nil {
			utils.EncodeError(g, err, g.Writer)
			return
		}
		g.JSON(http.StatusCreated, dataAccount)
	}
}

// DeleteAccount implements AccountEndpoint
func (a *accountHandler) DeleteAccount() gin.HandlerFunc {
	panic("unimplemented")
}

// ListAccount implements AccountEndpoint
func (a *accountHandler) ListAccount() gin.HandlerFunc {
	panic("unimplemented")
}

// UpdateAccount implements AccountEndpoint
func (a *accountHandler) UpdateAccount() gin.HandlerFunc {
	panic("unimplemented")
}

// UpdateMultiAccount implements AccountEndpoint
func (a *accountHandler) UpdateMultiAccount() gin.HandlerFunc {
	panic("unimplemented")
}

// ViewAccount implements AccountEndpoint
func (a *accountHandler) ViewAccount() gin.HandlerFunc {
	panic("unimplemented")
}

func NewAccountEndpoint(svc domains.Service) AccountEndpoint {
	return &accountHandler{service: svc}
}
