package middlewares

import (
	"context"
	"time"

	"github.com/bagasunix/ginclean/server/domains/entities"
	"github.com/bagasunix/ginclean/server/endpoints/requests"
	"github.com/bagasunix/ginclean/server/endpoints/responses"
	"go.uber.org/zap"
)

// CreateMultiRole implements domains.Service
func (l *loggingMiddleware) CreateMultiRole(ctx context.Context, req *[]requests.CreateRole) (res *responses.ListMultiple[entities.Role, entities.Role], err error) {
	defer func(begin time.Time) {
		l.logs.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "CreateRole"), zap.Any("response", string(res.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.CreateMultiRole(ctx, req)
}

// CreateRole implements domains.Service
func (l *loggingMiddleware) CreateRole(ctx context.Context, req *requests.CreateRole) (res *responses.EntityId, err error) {
	defer func(begin time.Time) {
		l.logs.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "CreateRole"), zap.Any("request", string(req.ToJSON())), zap.Any("response", string(res.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.CreateRole(ctx, req)
}

// DeleteRole implements domains.Service
func (l *loggingMiddleware) DeleteRole(ctx context.Context, req *requests.EntityId) (res *responses.Empty, err error) {
	defer func(begin time.Time) {
		l.logs.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "DeleteRole"), zap.Any("request", string(req.ToJSON())), zap.Any("response", string(res.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.DeleteRole(ctx, req)
}

// ListRole implements domains.Service
func (l *loggingMiddleware) ListRole(ctx context.Context, req *requests.BaseList) (res *responses.ListEntity[entities.Role], err error) {
	defer func(begin time.Time) {
		l.logs.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "ListRole"), zap.Any("request", string(req.ToJSON())), zap.Any("response", string(res.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.ListRole(ctx, req)
}

// UpdateMultipleRole implements domains.Service
func (l *loggingMiddleware) UpdateMultipleRole(ctx context.Context, req *[]requests.UpdateRole) (res *responses.ListMultiple[entities.Role, entities.Role], err error) {
	defer func(begin time.Time) {
		l.logs.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "UpdateRoleMulti"), zap.Any("response", string(res.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.UpdateMultipleRole(ctx, req)
}

// UpdateRole implements domains.Service
func (l *loggingMiddleware) UpdateRole(ctx context.Context, req *requests.UpdateRole) (res *responses.Empty, err error) {
	defer func(begin time.Time) {
		l.logs.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "UpdateRole"), zap.Any("request", string(req.ToJSON())), zap.Any("response", string(res.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.UpdateRole(ctx, req)
}

// ViewRole implements domains.Service
func (l *loggingMiddleware) ViewRole(ctx context.Context, req *requests.EntityId) (res *responses.ViewEntity[*entities.Role], err error) {
	defer func(begin time.Time) {
		l.logs.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "ViewRole"), zap.Any("request", string(req.ToJSON())), zap.Any("response", string(res.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.ViewRole(ctx, req)
}
