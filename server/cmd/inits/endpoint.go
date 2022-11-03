package inits

import (
	"github.com/bagasunix/ginclean/domains"
	"github.com/bagasunix/ginclean/endpoints"
)

func InitEndpoints(svc domains.Service) endpoints.Endpoints {
	a := endpoints.NewEndpointsBuilder()
	a.SetService(svc)
	return *a.Build()
}
