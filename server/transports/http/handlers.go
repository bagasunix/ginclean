package http

import (
	"github.com/bagasunix/ginclean/server/endpoints"
	"github.com/bagasunix/ginclean/server/endpoints/middlewares"
	"github.com/bagasunix/ginclean/server/transports/http/handlers"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewHttpHandler(logs *zap.Logger, eps endpoints.Endpoints) *gin.Engine {
	r := gin.New()
	gin.SetMode(gin.DebugMode)
	r.Use(ginzap.RecoveryWithZap(logs, true))
	r.Use(middlewares.GinContextToContextMiddleware())
	r.Use(middlewares.CORSMiddleware())

	// Create an account group
	handlers.MakeRoleHandler(logs, eps.RoleEndpoint, r.Group("/v1/role"))
	handlers.MakeUserHandler(logs, eps.UserEndpoint, r.Group("/v1/user"))
	handlers.MakeAuthHandler(eps.RefreshToken, r.Group("/v1/auth"))
	return r
}
