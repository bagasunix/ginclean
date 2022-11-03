package inits

import (
	"github.com/bagasunix/ginclean/server/domains"
	"github.com/bagasunix/ginclean/server/domains/data/repositories"
)

func InitService(repositories repositories.Repositories) domains.Service {
	svc := domains.NewServiceBuilder(repositories)
	return svc.Build()
}
