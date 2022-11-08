package endpoints

import (
	"net/http"

	"github.com/bagasunix/ginclean/server/domains"
	"github.com/bagasunix/ginclean/server/endpoints/requests"
	"github.com/bagasunix/ginclean/server/endpoints/utils"
	"github.com/gin-gonic/gin"
)

type UserEndpoint interface {
	CreateUser() gin.HandlerFunc
}

type userHandler struct {
	service domains.Service
}

// CreateUser implements UserEndpoint
func (u *userHandler) CreateUser() gin.HandlerFunc {
	return func(g *gin.Context) {
		var req requests.CreateAccount
		if err := g.Bind(&req); err != nil {
			utils.EncodeError(g, err, g.Writer)
			return
		}
		dataRole, err := u.service.CreateAccount(g, &req)
		if err != nil {
			utils.EncodeError(g, err, g.Writer)
			return
		}
		g.JSON(http.StatusCreated, dataRole)
	}
}

func NewUserEndpoint(svc domains.Service) UserEndpoint {
	return &userHandler{service: svc}
}