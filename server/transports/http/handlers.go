package http

import (
	"github.com/bagasunix/ginclean/endpoints"
	"github.com/bagasunix/ginclean/transports/http/handlers"
	"github.com/gin-gonic/gin"
)

func NewHttpHandler(r *gin.Engine, eps endpoints.Endpoints) *gin.Engine {

	// Create an account group
	handlers.MakeUserHandler(eps.RoleEndpoint, r.Group("/user"))
	return r
}
