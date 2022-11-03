package inits

import (
	"net/http"

	"github.com/bagasunix/ginclean/server/endpoints"
	transportHttp "github.com/bagasunix/ginclean/server/transports/http"
	"github.com/gin-gonic/gin"
)

func InitHttpHandler(endpoints endpoints.Endpoints) http.Handler {
	r := gin.Default()
	return transportHttp.NewHttpHandler(r, endpoints)
}
