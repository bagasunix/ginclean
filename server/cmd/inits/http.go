package inits

import (
	"net/http"

	"github.com/bagasunix/ginclean/server/endpoints"
	transportHttp "github.com/bagasunix/ginclean/server/transports/http"

	"go.uber.org/zap"
)

func InitHttpHandler(logs *zap.Logger, endpoints endpoints.Endpoints) http.Handler {

	return transportHttp.NewHttpHandler(logs, endpoints)
}
