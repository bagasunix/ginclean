package inits

import (
	"net/http"

	"github.com/bagasunix/ginclean/server/endpoints"
	transportHttp "github.com/bagasunix/ginclean/server/transports/http"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitHttpHandler(logs zap.Logger, endpoints endpoints.Endpoints) http.Handler {
	// r := gin.Default()
	r := gin.New()

	return transportHttp.NewHttpHandler(logs, r, endpoints)
}
