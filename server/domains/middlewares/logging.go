package middlewares

import (
	"github.com/bagasunix/ginclean/server/domains"
	"go.uber.org/zap"
)

type loggingMiddleware struct {
	logs zap.Logger
	next domains.Service
}

func LoggingMiddleware(logs zap.Logger) domains.Middleware {
	return func(next domains.Service) domains.Service {
		return &loggingMiddleware{logs: logs, next: next}
	}
}
