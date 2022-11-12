package endpoints

import (
	"github.com/bagasunix/ginclean/server/domains"
	"github.com/gin-gonic/gin"
)

type Endpoints struct {
	RoleEndpoint RoleEndpoint
	UserEndpoint UserEndpoint
}

// Builder Object for Endpoints
type EndpointsBuilder struct {
	service domains.Service
	// mdw     map[string][]Middleware
	mdw map[string]gin.HandlersChain
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
	o.UserEndpoint = NewUserEndpoint(b.service)
	return o
}

// Setter method for the field service of type domains.Service in the object EndpointsBuilder
func (e *EndpointsBuilder) SetService(service domains.Service) {
	e.service = service
}

// Setter method for the field mds of type gin.HandlerFunc in the object EndpointsBuilder
func (e *EndpointsBuilder) SetMdw(mdw map[string]gin.HandlersChain) {
	e.mdw = mdw
}
