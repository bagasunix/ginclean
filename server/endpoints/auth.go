package endpoints

import (
	"net/http"

	"github.com/bagasunix/ginclean/server/domains"
	"github.com/bagasunix/ginclean/server/domains/entities"
	"github.com/bagasunix/ginclean/server/endpoints/requests"
	"github.com/bagasunix/ginclean/server/endpoints/utils"
	"github.com/gin-gonic/gin"
)

type AuthEndpoint interface {
	RefreshToken() gin.HandlerFunc
}

type authHandler struct {
	service domains.Service
}

// RefreshToken implements AuthEndpoint
func (a *authHandler) RefreshToken() gin.HandlerFunc {
	return func(g *gin.Context) {
		reqBuild := entities.NewClientBuilder()
		reqBuild.SetIpClient(g.ClientIP())
		reqBuild.SetUserAgent(g.Request.UserAgent())
		g.Set("clients", reqBuild.Build())

		var req requests.Token
		if err := g.Bind(&req); err != nil {
			utils.EncodeError(g, err, g.Writer)
			return
		}
		dataToken, err := a.service.RefreshToken(g, &req)
		if err != nil {
			utils.EncodeError(g, err, g.Writer)
			return
		}

		g.JSON(http.StatusCreated, dataToken)

	}
}

func NewAuthEndpoint(svc domains.Service) AuthEndpoint {
	return &authHandler{service: svc}
}
