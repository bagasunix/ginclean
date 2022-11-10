package inits

import (
	"net/http"

	"github.com/bagasunix/ginclean/server/endpoints"
	"github.com/bagasunix/ginclean/server/endpoints/middlewares"
	transportHttp "github.com/bagasunix/ginclean/server/transports/http"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

func InitHttpHandler(logs *zap.Logger, endpoints endpoints.Endpoints) http.Handler {
	r := gin.New()
	gin.SetMode(gin.DebugMode)
	r.Use(ginzap.RecoveryWithZap(logs, true))
	r.Use(middlewares.GinContextToContextMiddleware())
	r.Use(middlewares.CORSMiddleware())

	return transportHttp.NewHttpHandler(logs, r, endpoints)
}
