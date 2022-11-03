package endpoints

import (
	"github.com/bagasunix/ginclean/domains"
)

type Endpoints struct {
	RoleEndpoint RoleEndpoint
}

// Builder Object for Endpoints
type EndpointsBuilder struct {
	service domains.Service
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
