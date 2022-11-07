package middlewares

import (
	"context"
	"time"

	"github.com/bagasunix/ginclean/server/domains/entities"
	"github.com/bagasunix/ginclean/server/endpoints/requests"
	"github.com/bagasunix/ginclean/server/endpoints/responses"
	"go.uber.org/zap"
)

// CreateAccount implements domains.Service
func (l *loggingMiddleware) CreateAccount(ctx context.Context, req *requests.CreateAccount) (res *responses.EntityId, err error) {
	defer func(begin time.Time) {
		l.logs.Log(zap.DebugLevel, "Middleware Domain", zap.String("method", "CreateAccount"), zap.Any("request", string(req.ToJSON())), zap.Any("response", string(res.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.CreateAccount(ctx, req)
}

// DeleteAccount implements domains.Service
func (l *loggingMiddleware) DeleteAccount(ctx context.Context, req *requests.EntityId) (res *responses.Empty, err error) {
	panic("unimplemented")
}

// DisableAccount implements domains.Service
func (l *loggingMiddleware) DisableAccount(ctx context.Context, request *requests.EntityId) (response *responses.Empty, err error) {
	panic("unimplemented")
}

// DisableMultipleAccount implements domains.Service
func (l *loggingMiddleware) DisableMultipleAccount(ctx context.Context, req []string) (res *responses.ListMultiple[*entities.Account, *entities.Account], err error) {
	panic("unimplemented")
}

// ListAccount implements domains.Service
func (l *loggingMiddleware) ListAccount(ctx context.Context, req *requests.BaseList) (res *responses.ListEntity[entities.Account], err error) {
	panic("unimplemented")
}

// ViewAccountByEmail implements domains.Service
func (l *loggingMiddleware) ViewAccountByEmail(ctx context.Context, email string) (res *responses.ViewEntity[*entities.Account], err error) {
	panic("unimplemented")
}

// ViewAccountByID implements domains.Service
func (l *loggingMiddleware) ViewAccountByID(ctx context.Context, req *requests.EntityId) (res *responses.ViewEntity[*entities.Account], err error) {
	panic("unimplemented")
}
