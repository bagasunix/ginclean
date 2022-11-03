package http

import (
	"github.com/bagasunix/ginclean/server/endpoints"
	"github.com/bagasunix/ginclean/server/endpoints/middlewares"
	"github.com/bagasunix/ginclean/server/transports/http/handlers"
	"github.com/gin-gonic/gin"
)

func NewHttpHandler(r *gin.Engine, eps endpoints.Endpoints) *gin.Engine {
	r.Use(middlewares.CORSMiddleware())

	// Create an account group
	handlers.MakeUserHandler(eps.RoleEndpoint, r.Group("/v1/role"))
	return r
}
