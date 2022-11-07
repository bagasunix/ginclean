package middlewares

import (
	"context"
	"time"

	"github.com/bagasunix/ginclean/server/endpoints"
	"go.uber.org/zap"
)

func Logging(logs zap.Logger) endpoints.Middleware {
	return func(e endpoints.Endpoint) endpoints.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logs.Log(zap.DebugLevel, "Middleware Endpoints", zap.Any("transport_error", err), zap.Any("took", time.Since(begin)))
			}(time.Now())
			return e(ctx, request)
		}
	}
}
