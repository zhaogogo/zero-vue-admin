package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	oteltrace "go.opentelemetry.io/otel/trace"
	"sync"

	"golang.org/x/sync/errgroup"
)

var (
	cacheUsercenterRoleMenuRoleIdPrefix = "cache:usercenter:roleMenu:role_id:"
)

var _ RoleMenuModel = (*customRoleMenuModel)(nil)

type (
	// RoleMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoleMenuModel.
	RoleMenuModel interface {
		roleMenuModel
		FindByRoleID(ctx context.Context, id []uint64) (menuIDs []uint64, err error)
	}

	customRoleMenuModel struct {
		*defaultRoleMenuModel
	}
)

// NewRoleMenuModel returns a model for the database table.
func NewRoleMenuModel(conn sqlx.SqlConn, c cache.CacheConf) RoleMenuModel {
	return &customRoleMenuModel{
		defaultRoleMenuModel: newRoleMenuModel(conn, c),
	}
}
func (m *defaultRoleMenuModel) FindByRoleID(ctx context.Context, roleIDs []uint64) (menuIDs []uint64, err error) {
	var (
		sqlqueryRoleIDs  []uint64
		cacheResRoleMenu []RoleMenu
	)
	tracer := otel.GetTracerProvider().Tracer("go-zero")
	redisCtx, _ := context.WithCancel(ctx)
	gRedis, _ := errgroup.WithContext(redisCtx)
	for _, roleid := range roleIDs {
		v := roleid
		gRedis.Go(func() error {
			var (
				roleMenu = []RoleMenu{}
				span     oteltrace.Span
				ctx      context.Context
			)

			usercenterRoleMenuRoleIdKey := fmt.Sprintf("%s%v", cacheUsercenterRoleMenuRoleIdPrefix, v)

			ctx, span = tracer.Start(redisCtx, "redis.cmds_describe", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
			span.SetAttributes(attribute.String("redis.cmds", fmt.Sprintf("get %s", usercenterRoleMenuRoleIdKey)))
			defer span.End()

			err := m.GetCacheCtx(ctx, usercenterRoleMenuRoleIdKey, &roleMenu)
			if err == sql.ErrNoRows {
				//如果缓存中是 * 什么也不做
				span.SetStatus(codes.Unset, fmt.Sprintf("get %s", usercenterRoleMenuRoleIdKey))
				sqlqueryRoleIDs = append(sqlqueryRoleIDs, v)
			} else if err != nil {
				// redis查询错误，加入sql查询列表
				span.SetStatus(codes.Error, fmt.Sprintf("get %s", usercenterRoleMenuRoleIdKey))
				logx.Errorf("redis查询错误，表：%v, role_id: %v", m.table, v)
				sqlqueryRoleIDs = append(sqlqueryRoleIDs, v)
			} else {
				span.SetStatus(codes.Ok, fmt.Sprintf("get %s", usercenterRoleMenuRoleIdKey))
				cacheResRoleMenu = append(cacheResRoleMenu, roleMenu...)
			}
			return nil
		})
	}
	gRedis.Wait()
	var (
		sqlqueryResp = make(map[uint64][]RoleMenu)
		mu           sync.Mutex
	)
	sqlctx, _ := context.WithCancel(ctx)
	gSql, _ := errgroup.WithContext(sqlctx)
	//fmt.Println("sqlqueryRoleIDs ==>", sqlqueryRoleIDs)
	if len(sqlqueryRoleIDs) != 0 {
		tracer := otel.GetTracerProvider().Tracer("go-zero")
		for _, roleID := range sqlqueryRoleIDs {
			roleid := roleID
			var res = []RoleMenu{}
			query := fmt.Sprintf("SELECT %s FROM %s WHERE `role_id` = \"%v\"", roleMenuRows, m.table, roleid)

			gSql.Go(func() error {
				var span oteltrace.Span
				ctx, span = tracer.Start(sqlctx, "sql.method_describe", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
				span.SetAttributes(attribute.String("sql", query))
				defer span.End()
				//fmt.Println("sqlqueryRoleIDs: ---->", sqlqueryRoleIDs, query)
				err = m.QueryRowsNoCacheCtx(ctx, &res, query)
				fmt.Println("====>", query, err)
				if err != nil {
					span.SetStatus(codes.Error, query)
					span.SetAttributes(attribute.String("error", err.Error()))
					logx.Errorf("数据库查询错误, table: %s, role_id: %v", m.table, roleid)
					return err
				} else {
					span.SetStatus(codes.Ok, query)
					m.SetCacheCtx(ctx, fmt.Sprintf("%s%v", cacheUsercenterRoleMenuRoleIdPrefix, roleid), res)
					//fmt.Println("set ", fmt.Sprintf("%s%v", cacheUsercenterRoleMenuRoleIdPrefix, roleid), e, res)
					mu.Lock()
					sqlqueryResp[roleid] = append(sqlqueryResp[roleid], res...)
					mu.Unlock()
				}
				return nil
			})

		}
	}
	err = gSql.Wait()
	//fmt.Printf("roleMenu ===> %v res: %#v\n", err, sqlqueryResp)
	switch err {
	case nil:
		var (
			menus    []uint64
			menusSet = make(map[uint64]int)
		)
		for _, v := range cacheResRoleMenu {
			menusSet[v.MenuId] += 1
		}

		for _, v := range sqlqueryResp {
			for _, v2 := range v {
				menusSet[v2.MenuId] += 1
			}
		}
		for k, _ := range menusSet {
			menus = append(menus, k)
		}
		return menus, nil
	case sqlx.ErrNotFound:
		logx.Errorf("sql查询无数据，表: %v, role_id: %v, 返回缓存数据", m.table, sqlqueryRoleIDs)
		var (
			menus []uint64
		)
		for _, v := range cacheResRoleMenu {
			menus = append(menus, v.MenuId)
		}
		if len(menus) == 0 {
			return nil, ErrNotFound
		}
		return menus, nil
	default:
		return nil, err
	}
}
