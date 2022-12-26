package system

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"
	"strings"
)

var (
	cacheChaosSystemUserMenuParamsUserIdPrefix = "cache:chaosSystem:userMenuParams:user_id:"
)
var _ UserMenuParamsModel = (*customUserMenuParamsModel)(nil)

type (
	// UserMenuParamsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserMenuParamsModel.
	UserMenuParamsModel interface {
		FindAll_NC(ctx context.Context) ([]UserMenuParams, error)
		Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error

		FindByUserID(ctx context.Context, redis *redis.Redis, userid uint64) ([]UserMenuParams, error)

		InsertMultiple(ctx context.Context, data []*UserMenuParams) (sql.Result, error)

		DeleteByMenuID(ctx context.Context, menuid uint64) error

		TransDeleteByMenuID(ctx context.Context, session sqlx.Session, menuid uint64) error
		TransDeleteByUserID(ctx context.Context, session sqlx.Session, userid uint64) error
		TransInsertMultiple(ctx context.Context, session sqlx.Session, data []*UserMenuParams) (sql.Result, error)
		TransDeleteNotINANDMenuID(ctx context.Context, session sqlx.Session, data []*UserMenuParams, menuid uint64) error
		TransUpdate(ctx context.Context, session sqlx.Session, data *UserMenuParams) error
	}

	customUserMenuParamsModel struct {
		*defaultUserMenuParamsModel
	}
)

// NewUserMenuParamsModel returns a model for the database table.
func NewUserMenuParamsModel(conn sqlx.SqlConn, c cache.CacheConf) UserMenuParamsModel {
	return &customUserMenuParamsModel{
		defaultUserMenuParamsModel: newUserMenuParamsModel(conn, c),
	}
}

func (m *defaultUserMenuParamsModel) DeleteByMenuID(ctx context.Context, menuid uint64) error {
	//寻找删除缓存key
	var (
		cacheAllKey     = []string{}
		usermenuparamss []UserMenuParams
	)
	query := fmt.Sprintf("SELECT %s FROM %s where `menu_id` = ?", userMenuParamsRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &usermenuparamss, query, menuid)
	if err != nil {
		return err
	}
	if len(usermenuparamss) != 0 {
		for _, usermenuparams := range usermenuparamss {
			cacheAllKey = append(cacheAllKey, fmt.Sprintf("%s%v", cacheChaosSystemUserMenuParamsUserIdPrefix, usermenuparams.UserId))
		}
	}

	//执行删除（数据库删除、缓存删除）
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `menu_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, menuid)
	}, cacheAllKey...)
	return err
}

func (m *defaultUserMenuParamsModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultUserMenuParamsModel) TransInsertMultiple(ctx context.Context, session sqlx.Session, data []*UserMenuParams) (sql.Result, error) {
	chaosSystemUserMenuParamsUserIdKeys := []string{}
	values := []string{}
	useridSet := map[uint64]int64{}
	query := fmt.Sprintf("insert into %s (%s) values ", m.table, userMenuParamsRowsExpectAutoSet)
	for _, v := range data {
		useridSet[v.UserId] += 1
		values = append(values, fmt.Sprintf("(%v, %v, \"%v\", \"%v\", \"%v\")", v.UserId, v.MenuId, v.Type, v.Key, v.Value))
	}
	for userid, _ := range useridSet {
		chaosSystemUserMenuParamsUserIdKeys = append(
			chaosSystemUserMenuParamsUserIdKeys,
			fmt.Sprintf("%s%v", cacheChaosSystemUserMenuParamsUserIdPrefix, userid),
		)
	}
	query += strings.Join(values, ",")
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		return session.ExecCtx(ctx, query)
	}, chaosSystemUserMenuParamsUserIdKeys...)
	return ret, err
}

func (m *defaultUserMenuParamsModel) TransUpdate(ctx context.Context, session sqlx.Session, data *UserMenuParams) error {
	chaosSystemUserMenuParamsUserIdKey := fmt.Sprintf("%s%v", cacheChaosSystemUserMenuParamsUserIdPrefix, data.UserId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userMenuParamsRowsWithPlaceHolder)
		return session.ExecCtx(ctx, query, data.UserId, data.MenuId, data.Type, data.Key, data.Value, data.Id)
	}, chaosSystemUserMenuParamsUserIdKey)
	return err
}

func (m *defaultUserMenuParamsModel) TransDeleteByMenuID(ctx context.Context, session sqlx.Session, menuid uint64) error {
	//寻找删除缓存key
	var (
		cacheAllKey     = []string{}
		usermenuparamss []UserMenuParams
	)
	query := fmt.Sprintf("SELECT %s FROM %s where `menu_id` = ?", userMenuParamsRows, m.table)
	err := session.QueryRowsCtx(ctx, &usermenuparamss, query, menuid)
	if err != nil {
		return err
	}
	if len(usermenuparamss) != 0 {
		for _, usermenuparams := range usermenuparamss {
			cacheAllKey = append(cacheAllKey, fmt.Sprintf("%s%v", cacheChaosSystemUserMenuParamsUserIdPrefix, usermenuparams.UserId))
		}
	}

	//执行删除（数据库删除、缓存删除）
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `menu_id` = ?", m.table)
		return session.ExecCtx(ctx, query, menuid)
	}, cacheAllKey...)
	return err
}

func (m *defaultUserMenuParamsModel) TransDeleteByUserID(ctx context.Context, session sqlx.Session, userid uint64) error {
	chaosSystemUserMenuParamsUserIdKey := fmt.Sprintf("%s%v", cacheChaosSystemUserMenuParamsUserIdPrefix, userid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
		return session.ExecCtx(ctx, query, userid)
	}, chaosSystemUserMenuParamsUserIdKey)
	return err
}

func (m *defaultUserMenuParamsModel) FindByUserID(ctx context.Context, redis *redis.Redis, userID uint64) ([]UserMenuParams, error) {
	chaosSystemUserMenuParamsUserIdKey := fmt.Sprintf("%s%v", cacheChaosSystemUserMenuParamsUserIdPrefix, userID)
	var resp []UserMenuParams
	err := m.GetCacheCtx(ctx, chaosSystemUserMenuParamsUserIdKey, &resp)
	if err == nil {
		return resp, nil
	}
	if err.Error() == "placeholder" {
		return nil, ErrNotFound
	}
	if err == sql.ErrNoRows {
		query := fmt.Sprintf("select %s from %s where `user_id` = ?", userMenuParamsRows, m.table)
		err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userID)
		if err != nil {
			return nil, err
		}
		if len(resp) == 0 {
			err := redis.Setex(chaosSystemUserMenuParamsUserIdKey, "*", int(unstable.AroundDuration(cacheOption.NotFoundExpiry).Seconds()))
			if err != nil {
				logx.Errorf("设置缓存失败, key: %v, error: %v", chaosSystemUserMenuParamsUserIdKey, err)
			}
			return nil, ErrNotFound
		}
		err = m.SetCacheCtx(ctx, chaosSystemUserMenuParamsUserIdKey, resp)
		if err != nil {
			logx.Errorf("设置缓存失败, key: %v, error: %v", chaosSystemUserMenuParamsUserIdKey, err)
		}
		return resp, nil
	}

	return nil, err
}

func (m *defaultUserMenuParamsModel) FindAll_NC(ctx context.Context) ([]UserMenuParams, error) {
	var resp []UserMenuParams
	query := fmt.Sprintf("select %s from %s", userMenuParamsRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, ErrNotFound
	}

	return resp, nil
}

//data列表中的menuid 要和传入的menuid一致
func (m *defaultUserMenuParamsModel) TransDeleteNotINANDMenuID(ctx context.Context, session sqlx.Session, data []*UserMenuParams, menuid uint64) error {
	//查找删除缓存Key
	useridSet := make(map[uint64]int64)
	ids := []string{}
	var cacheAllKey = []string{}
	for _, v := range data {
		if v.MenuId != menuid {
			return errors.New("更新数据Slice MenuId字段和传入的menuid不一致")
		}
		ids = append(ids, strconv.FormatUint(v.Id, 10))
	}
	var usermenuparamssDel []UserMenuParams
	query := fmt.Sprintf("SELECT %s FROM %s where `menu_id` = ? AND `id` NOT IN (%s)", userMenuParamsRows, m.table, strings.Join(ids, ","))
	err := session.QueryRowsCtx(ctx, &usermenuparamssDel, query, menuid)
	if err != nil {
		return err
	}
	if len(usermenuparamssDel) != 0 {
		//删除的行
		for _, usermenuparams := range usermenuparamssDel {
			useridSet[usermenuparams.UserId] += 1
			cacheAllKey = append(cacheAllKey, fmt.Sprintf("%s%v", cacheChaosSystemUserMenuParamsUserIdPrefix, usermenuparams.UserId))
		}
	}
	//数据的用户改变了 查找老的用户 删除缓存
	var usermenuparamssOld []UserMenuParams
	query = fmt.Sprintf("SELECT %s FROM %s where `menu_id` = ? AND `id` IN (%s)", userMenuParamsRows, m.table, strings.Join(ids, ","))
	err = session.QueryRowsCtx(ctx, &usermenuparamssOld, query, menuid)
	if err != nil {
		return err
	}
	if len(usermenuparamssOld) != 0 {
		for _, usermenuparams := range usermenuparamssOld {
			useridSet[usermenuparams.UserId] += 1
		}
	}

	for userid, _ := range useridSet {
		cacheAllKey = append(cacheAllKey, fmt.Sprintf("%s%v", cacheChaosSystemUserMenuParamsUserIdPrefix, userid))
	}

	// shanchu
	query = fmt.Sprintf("DELETE FROM %s WHERE `id` NOT IN (%s) AND `menu_id` = ?", m.table, strings.Join(ids, ","))
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		return session.ExecCtx(ctx, query, menuid)
	}, cacheAllKey...)
	if err != nil {
		return err
	}
	return nil
}

func (m *defaultUserMenuParamsModel) InsertMultiple(ctx context.Context, data []*UserMenuParams) (sql.Result, error) {
	chaosSystemUserMenuParamsUserIdKeys := []string{}
	values := []string{}
	useridSet := map[uint64]int64{}
	query := fmt.Sprintf("insert into %s (%s) values ", m.table, userMenuParamsRowsExpectAutoSet)
	for _, v := range data {
		useridSet[v.UserId] += 1
		values = append(values, fmt.Sprintf("(%v, %v, \"%v\", \"%v\", \"%v\")", v.UserId, v.MenuId, v.Type, v.Key, v.Value))
	}
	for userid, _ := range useridSet {
		chaosSystemUserMenuParamsUserIdKeys = append(
			chaosSystemUserMenuParamsUserIdKeys,
			fmt.Sprintf("%s%v", cacheChaosSystemUserMenuParamsUserIdPrefix, userid),
		)
	}
	query += strings.Join(values, ",")
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		return conn.ExecCtx(ctx, query)
	}, chaosSystemUserMenuParamsUserIdKeys...)
	return ret, err
}
