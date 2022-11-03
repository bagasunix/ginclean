package inits

import (
	"github.com/bagasunix/ginclean/server/domains"
	"github.com/bagasunix/ginclean/server/endpoints"
)

func InitEndpoints(svc domains.Service) endpoints.Endpoints {
	a := endpoints.NewEndpointsBuilder()
	a.SetService(svc)
	return *a.Build()
}
