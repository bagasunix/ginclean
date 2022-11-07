package domains

import (
	"github.com/bagasunix/ginclean/server/domains/data/repositories"
	"github.com/bagasunix/ginclean/server/domains/usecases"
	"go.uber.org/zap"
)

type Service interface {
	usecases.RoleService
}

type service struct {
	usecases.RoleService
}

type ServiceBuilder struct {
	logs       zap.Logger
	repo       repositories.Repositories
	middleware []Middleware
}

func NewServiceBuilder(logs zap.Logger, repo repositories.Repositories) *ServiceBuilder {
	s := new(ServiceBuilder)
	s.repo = repo
	s.logs = logs
	return s
}

func buildService(logs zap.Logger, repo repositories.Repositories) Service {
	svc := new(service)
	svc.RoleService = usecases.NewRole(logs, repo)
	return svc
}

func (s *ServiceBuilder) Build() Service {
	svc := buildService(s.logs, s.repo)
	for _, v := range s.middleware {
		svc = v(svc)
	}
	return svc
}

// Setter method for the field middleware of type []Middleware in the object ServiceBuilder
func (s *ServiceBuilder) SetMiddleware(middleware []Middleware) {
	s.middleware = middleware
}
