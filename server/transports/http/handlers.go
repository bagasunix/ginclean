package http

import (
	"github.com/bagasunix/ginclean/server/endpoints"
	"github.com/bagasunix/ginclean/server/transports/http/handlers"
	"github.com/gin-gonic/gin"
)

func NewHttpHandler(r *gin.Engine, eps endpoints.Endpoints) *gin.Engine {

	// Create an account group
	handlers.MakeRoleHandler(eps.RoleEndpoint, r.Group("/v1/role"))
	handlers.MakeUserHandler(eps.UserEndpoint, r.Group("/v1/user"))
	return r
}
