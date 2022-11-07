package inits

import (
	"github.com/bagasunix/ginclean/server/domains"
	"github.com/bagasunix/ginclean/server/endpoints"
	"github.com/bagasunix/ginclean/server/endpoints/middlewares"
	"go.uber.org/zap"
)

func InitEndpoints(logs zap.Logger, svc domains.Service) endpoints.Endpoints {
	a := endpoints.NewEndpointsBuilder()
	a.SetMdw(getEndpointMiddleware(logs))
	a.SetService(svc)
	return *a.Build()
}

func getEndpointMiddleware(logs zap.Logger) (mw map[string][]endpoints.Middleware) {
	mw = map[string][]endpoints.Middleware{}
	addDefaultEndpointMiddleware(logs, mw)
	return mw
}

func middlewaresWithAuthentication(logs zap.Logger, method string) []endpoints.Middleware {
	mw := defaultMiddlewares(logs, method)
	return mw
	// return append(mw, middlewares.Authentication())
}

func defaultMiddlewares(logs zap.Logger, method string) []endpoints.Middleware {
	return []endpoints.Middleware{
		middlewares.Logging(*logs.With(zap.Any("method", method))),
	}
}

func addDefaultEndpointMiddleware(logs zap.Logger, mw map[string][]endpoints.Middleware) {
	mw[endpoints.CREATE_ROLE] = middlewaresWithAuthentication(logs, endpoints.CREATE_ROLE)
	mw[endpoints.LIST_ROLE] = middlewaresWithAuthentication(logs, endpoints.LIST_ROLE)
	mw[endpoints.VIEW_ROLE] = middlewaresWithAuthentication(logs, endpoints.VIEW_ROLE)
	mw[endpoints.UPDATE_ROLE] = middlewaresWithAuthentication(logs, endpoints.UPDATE_ROLE)
	mw[endpoints.DELETE_ROLE] = middlewaresWithAuthentication(logs, endpoints.DELETE_ROLE)
}
