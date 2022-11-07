package inits

import (
	"github.com/bagasunix/ginclean/server/domains"
	"github.com/bagasunix/ginclean/server/endpoints"
)

func InitEndpoints(svc domains.Service) endpoints.Endpoints {
	a := endpoints.NewEndpointsBuilder()
	a.SetService(svc)
	// a.SetMdw(getEndpointMiddleware())
	return *a.Build()
}

// func getEndpointMiddleware() (mw map[string][]endpoints.Middleware) {
// 	mw = map[string][]endpoints.Middleware{}
// 	addDefaultEndpointMiddleware(mw)
// 	return mw
// }

// func middlewaresWithAuthentication(method string) []endpoints.Middleware {
// 	mw := defaultMiddlewares(method)
// 	return mw
// 	// return append(mw, middlewares.Authentication())
// }

// func defaultMiddlewares(method string) []endpoints.Middleware {
// 	return []endpoints.Middleware{}
// }

// func addDefaultEndpointMiddleware(mw map[string][]endpoints.Middleware) {
// 	mw[endpoints.CREATE_ROLE] = middlewaresWithAuthentication(endpoints.CREATE_ROLE)
// 	mw[endpoints.LIST_ROLE] = middlewaresWithAuthentication(endpoints.LIST_ROLE)
// 	mw[endpoints.VIEW_ROLE] = middlewaresWithAuthentication(endpoints.VIEW_ROLE)
// 	mw[endpoints.UPDATE_ROLE] = middlewaresWithAuthentication(endpoints.UPDATE_ROLE)
// 	mw[endpoints.DELETE_ROLE] = middlewaresWithAuthentication(endpoints.DELETE_ROLE)
// }
