package inits

import (
	"github.com/bagasunix/ginclean/server/domains"
	"github.com/bagasunix/ginclean/server/endpoints"
	"github.com/bagasunix/ginclean/server/endpoints/middlewares"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitEndpoints(logs zap.Logger, svc domains.Service) endpoints.Endpoints {
	a := endpoints.NewEndpointsBuilder()
	a.SetMdw(getEndpointMiddleware(logs))
	a.SetService(svc)
	return *a.Build()
}

func getEndpointMiddleware(logs zap.Logger) (mw map[string]gin.HandlersChain) {
	mw = map[string]gin.HandlersChain{}
	addDefaultEndpointMiddleware(logs, mw)
	return mw
}

func middlewaresWithAuthentication(logs zap.Logger, method string) gin.HandlersChain {
	mw := defaultMiddlewares(logs, method)
	// return mw
	mw = append(mw, middlewares.CORSMiddleware())
	mw = append(mw, middlewares.GinContextToContextMiddleware())
	mw = append(mw, ginzap.RecoveryWithZap(&logs, true))
	return mw
}

func defaultMiddlewares(logs zap.Logger, method string) gin.HandlersChain {
	return gin.HandlersChain{
		// middlewares.Logging(*logs.With(zap.Any("method", method))),
	}
}

func addDefaultEndpointMiddleware(logs zap.Logger, mw map[string]gin.HandlersChain) {
	mw[endpoints.CREATE_ROLE] = middlewaresWithAuthentication(logs, endpoints.CREATE_ROLE)
	mw[endpoints.LIST_ROLE] = middlewaresWithAuthentication(logs, endpoints.LIST_ROLE)
	mw[endpoints.VIEW_ROLE] = middlewaresWithAuthentication(logs, endpoints.VIEW_ROLE)
	mw[endpoints.UPDATE_ROLE] = middlewaresWithAuthentication(logs, endpoints.UPDATE_ROLE)
	mw[endpoints.DELETE_ROLE] = middlewaresWithAuthentication(logs, endpoints.DELETE_ROLE)
}
