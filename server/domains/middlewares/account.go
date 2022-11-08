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
		l.logs.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "CreateAccount"), zap.Any("request", string(req.ToJSON())), zap.Any("response", string(res.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.CreateAccount(ctx, req)
}

// DeleteAccount implements domains.Service
func (l *loggingMiddleware) DeleteAccount(ctx context.Context, req *requests.EntityId) (res *responses.Empty, err error) {
	defer func(begin time.Time) {
		l.logs.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "DeleteAccount"), zap.Any("request", string(req.ToJSON())), zap.Any("response", string(res.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.DeleteAccount(ctx, req)
}

// DisableAccount implements domains.Service
func (l *loggingMiddleware) DisableAccount(ctx context.Context, request *requests.EntityId) (response *responses.Empty, err error) {
	defer func(begin time.Time) {
		l.logs.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "DisableAccount"), zap.Any("request", string(request.ToJSON())), zap.Any("response", string(response.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.DisableAccount(ctx, request)
}

// DisableMultipleAccount implements domains.Service
func (l *loggingMiddleware) DisableMultipleAccount(ctx context.Context, req []string) (res *responses.ListMultiple[*entities.Account, *entities.Account], err error) {
	defer func(begin time.Time) {
		l.logs.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "DisableAccount"), zap.Any("request", req), zap.Any("response_success", res.Data.Success), zap.Any("response_failed", res.Data.Failed), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.DisableMultipleAccount(ctx, req)
}

// ListAccount implements domains.Service
func (l *loggingMiddleware) ListAccount(ctx context.Context, req *requests.BaseList) (res *responses.ListEntity[entities.Account], err error) {
	defer func(begin time.Time) {
		l.logs.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "ListAccount"), zap.Any("request", string(req.ToJSON())), zap.Any("response", string(res.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.ListAccount(ctx, req)
}

// ViewAccountByEmail implements domains.Service
func (l *loggingMiddleware) ViewAccountByEmail(ctx context.Context, email string) (res *responses.ViewEntity[*entities.Account], err error) {
	defer func(begin time.Time) {
		l.logs.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "ViewAccountByEmail"), zap.Any("request", string(email)), zap.Any("response", string(res.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.ViewAccountByEmail(ctx, email)
}

// ViewAccountByID implements domains.Service
func (l *loggingMiddleware) ViewAccountByID(ctx context.Context, req *requests.EntityId) (res *responses.ViewEntity[*entities.Account], err error) {
	defer func(begin time.Time) {
		l.logs.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "ViewAccountByID"), zap.Any("request", string(req.ToJSON())), zap.Any("response", string(res.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.ViewAccountByID(ctx, req)
}
