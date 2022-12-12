package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	oteltrace "go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
	"sync"
)

var _ MenuModel = (*customMenuModel)(nil)

type (
	// MenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMenuModel.
	MenuModel interface {
		menuModel
		FindByIDs(ctx context.Context, IDs []uint64) ([]Menu, error)
		FindAll(ctx context.Context) ([]Menu, error)
	}

	customMenuModel struct {
		*defaultMenuModel
	}
)

// NewMenuModel returns a model for the database table.
func NewMenuModel(conn sqlx.SqlConn, c cache.CacheConf) MenuModel {
	return &customMenuModel{
		defaultMenuModel: newMenuModel(conn, c),
	}
}

func (m *defaultMenuModel) FindByIDs(ctx context.Context, IDs []uint64) ([]Menu, error) {
	var (
		err             error
		mu              sync.Mutex
		sqlqueryMenuIDs []uint64

		cacheMenus []Menu
	)
	tracer := otel.GetTracerProvider().Tracer("go-zero")
	redisCtx, _ := context.WithCancel(ctx)
	gRedis, _ := errgroup.WithContext(redisCtx)
	for _, v := range IDs {
		id := v
		gRedis.Go(func() error {
			var menu Menu
			usercenterMenuIdKey := fmt.Sprintf("%s%v", cacheUsercenterMenuIdPrefix, id)
			//链路追踪
			ctx, span := tracer.Start(redisCtx, "redis.cmds_describe", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
			span.SetAttributes(attribute.String("redis.cmds", fmt.Sprintf("get %s", usercenterMenuIdKey)))
			defer span.End()

			if err := m.GetCacheCtx(ctx, usercenterMenuIdKey, &menu); err == sql.ErrNoRows {
				//如果缓存中是 * 什么也不做
				span.SetStatus(codes.Unset, fmt.Sprintf("get %s", usercenterMenuIdKey))
				mu.Lock()
				sqlqueryMenuIDs = append(sqlqueryMenuIDs, id)
				mu.Unlock()
			} else if err != nil {
				//redis查询错误
				span.SetStatus(codes.Error, fmt.Sprintf("get %s", usercenterMenuIdKey))
				logx.Errorf("redis查询错误，表: %v, id: %v", m.table, id)
				mu.Lock()
				sqlqueryMenuIDs = append(sqlqueryMenuIDs, id)
				mu.Unlock()
			} else {
				span.SetStatus(codes.Ok, fmt.Sprintf("get %s", usercenterMenuIdKey))
				mu.Lock()
				cacheMenus = append(cacheMenus, menu)
				mu.Unlock()
			}

			return nil
		})
	}
	gRedis.Wait()

	sqlctx, _ := context.WithCancel(ctx)
	gSql, _ := errgroup.WithContext(sqlctx)

	if len(sqlqueryMenuIDs) != 0 {
		tracer := otel.GetTracerProvider().Tracer("go-zero")
		for _, ID := range sqlqueryMenuIDs {
			id := ID
			var resp Menu
			query := fmt.Sprintf("SELECT %s FROM %s WHERE `id` = \"%v\"", menuRows, m.table, id)

			gSql.Go(func() error {
				ctx, span := tracer.Start(sqlctx, "sql.method_describe", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
				span.SetAttributes(attribute.String("sql", query))
				defer span.End()
				//fmt.Println("sqlqueryMenuIDs ====>", sqlqueryMenuIDs, query)
				err := m.QueryRowNoCacheCtx(ctx, &resp, query)
				if err != nil {
					span.SetStatus(codes.Error, query)
					span.SetAttributes(attribute.String("error", err.Error()))
					return err
				}

				span.SetStatus(codes.Ok, query)
				m.SetCacheCtx(ctx, fmt.Sprintf("%s%v", cacheUsercenterMenuIdPrefix, id), resp)
				mu.Lock()
				cacheMenus = append(cacheMenus, resp)
				mu.Unlock()
				return nil

			})

		}

	}
	err = gSql.Wait()
	switch {
	case err == nil || err == sqlx.ErrNotFound:
		resp := []Menu{}
		for _, v := range cacheMenus {
			if !v.DeleteTime.Valid {
				resp = append(resp, v)
			}
		}
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultMenuModel) FindAll(ctx context.Context) ([]Menu, error) {
	var resp []Menu
	query := fmt.Sprintf("SELECT %s FROM %s", menuRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
