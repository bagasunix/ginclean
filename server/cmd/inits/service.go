package inits

import (
	"github.com/bagasunix/ginclean/server/domains"
	"github.com/bagasunix/ginclean/server/domains/data/repositories"
	"github.com/bagasunix/ginclean/server/domains/middlewares"
	"go.uber.org/zap"
)

func InitService(logs zap.Logger, repositories repositories.Repositories) domains.Service {
	svc := domains.NewServiceBuilder(logs, repositories)
	svc.SetMiddleware(getServiceMiddleware(logs))

	return svc.Build()
}

func getServiceMiddleware(logs zap.Logger) []domains.Middleware {
	var mw []domains.Middleware
	mw = addDefaultServiceMiddleware(logs, mw)
	return mw
}

func addDefaultServiceMiddleware(logs zap.Logger, mw []domains.Middleware) []domains.Middleware {
	return append(mw, middlewares.LoggingMiddleware(logs))
}
