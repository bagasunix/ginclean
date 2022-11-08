package http

import (
	"github.com/bagasunix/ginclean/server/endpoints"
	"github.com/bagasunix/ginclean/server/endpoints/middlewares"
	"github.com/bagasunix/ginclean/server/transports/http/handlers"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewHttpHandler(logs *zap.Logger, r *gin.Engine, eps endpoints.Endpoints) *gin.Engine {
	r.Use(middlewares.CORSMiddleware())
	r.Use(ginzap.RecoveryWithZap(logs, true))

	// Create an account group
	handlers.MakeRoleHandler(eps.RoleEndpoint, r.Group("/v1/role"))
	handlers.MakeUserHandler(eps.UserEndpoint, r.Group("/v1/user"))
	return r
}
