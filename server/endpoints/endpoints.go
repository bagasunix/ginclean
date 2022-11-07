package endpoints

import (
	"context"

	"github.com/bagasunix/ginclean/server/domains"
)

type Endpoint func(ctx context.Context, req interface{}) (res interface{}, err error)
type Middleware func(Endpoint) Endpoint

type Endpoints struct {
	RoleEndpoint RoleEndpoint
}

// Builder Object for Endpoints
type EndpointsBuilder struct {
	service domains.Service
	// mdws    map[string][]gin.HandlerFunc
	// mdw     map[string][]Middleware
}

// Constructor for EndpointsBuilder
func NewEndpointsBuilder() *EndpointsBuilder {
	o := new(EndpointsBuilder)
	return o
}

// Build Method which creates Endpoints
func (b *EndpointsBuilder) Build() *Endpoints {
	o := new(Endpoints)
	o.RoleEndpoint = NewRoleEndpoint(b.service)
	return o
}

// Setter method for the field service of type domains.Service in the object EndpointsBuilder
func (e *EndpointsBuilder) SetService(service domains.Service) {
	e.service = service
}
