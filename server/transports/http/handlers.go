package http

import (
	"github.com/bagasunix/ginclean/server/endpoints"
	"github.com/bagasunix/ginclean/server/transports/http/handlers"
	"github.com/gin-gonic/gin"
)

func NewHttpHandler(r *gin.Engine, eps endpoints.Endpoints) *gin.Engine {

	// Create an account group
	handlers.MakeUserHandler(eps.RoleEndpoint, r.Group("/user"))
	return r
}
