package domains

import (
	"github.com/bagasunix/ginclean/domains/data/repositories"
	"github.com/bagasunix/ginclean/domains/usecases"
)

type Service interface {
	usecases.RoleService
}

type service struct {
	usecases.RoleService
}

type ServiceBuilder struct {
	repo       repositories.Repositories
	middleware []Middleware
}

func NewServiceBuilder(repo repositories.Repositories) *ServiceBuilder {
	s := new(ServiceBuilder)
	s.repo = repo
	return s
}

func buildService(repo repositories.Repositories) Service {
	svc := new(service)
	svc.RoleService = usecases.NewRole(repo)
	return svc
}

func (s *ServiceBuilder) Build() Service {
	svc := buildService(s.repo)
	for _, v := range s.middleware {
		svc = v(svc)
	}
	return svc
}
