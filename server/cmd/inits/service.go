package inits

import (
	"github.com/bagasunix/ginclean/server/domains"
	"github.com/bagasunix/ginclean/server/domains/data/repositories"
	"go.uber.org/zap"
)

func InitService(logs zap.Logger, repositories repositories.Repositories) domains.Service {
	svc := domains.NewServiceBuilder(logs, repositories)
	return svc.Build()
}
