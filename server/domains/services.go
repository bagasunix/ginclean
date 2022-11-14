package domains

import (
	"github.com/bagasunix/ginclean/server/domains/data/repositories"
	"github.com/bagasunix/ginclean/server/domains/usecases"
	"go.uber.org/zap"
)

type Service interface {
	usecases.RoleService
	usecases.AccountService
}

type service struct {
	usecases.RoleService
	usecases.AccountService
}

type ServiceBuilder struct {
	logs          *zap.Logger
	repo          repositories.Repositories
	middleware    []Middleware
	jwtKey        string
	jwtKeyRefresh string
}

func NewServiceBuilder(logs *zap.Logger, jwtKey string, jwtKeyRefresh string, repo repositories.Repositories) *ServiceBuilder {
	s := new(ServiceBuilder)
	s.jwtKey = jwtKey
	s.jwtKeyRefresh = jwtKeyRefresh
	s.logs = logs
	s.repo = repo
	return s
}

func buildService(logs *zap.Logger, jwtKey string, jwtKeyRefresh string, repo repositories.Repositories) Service {
	svc := new(service)
	svc.RoleService = usecases.NewRole(logs, repo)
	svc.AccountService = usecases.NewAccount(logs, jwtKey, jwtKeyRefresh, repo)
	return svc
}

func (s *ServiceBuilder) Build() Service {
	svc := buildService(s.logs, s.jwtKey, s.jwtKeyRefresh, s.repo)
	for _, v := range s.middleware {
		svc = v(svc)
	}
	return svc
}

// Setter method for the field middleware of type []Middleware in the object ServiceBuilder
func (s *ServiceBuilder) SetMiddleware(middleware []Middleware) {
	s.middleware = middleware
}
