package inits

import (
	"net/http"

	"github.com/bagasunix/ginclean/endpoints"
	transportHttp "github.com/bagasunix/ginclean/transports/http"
	"github.com/gin-gonic/gin"
)

func InitHttpHandler(endpoints endpoints.Endpoints) http.Handler {
	r := gin.Default()
	return transportHttp.NewHttpHandler(r, endpoints)
}
