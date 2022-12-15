package system

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserPageSetModel = (*customUserPageSetModel)(nil)

type (
	// UserPageSetModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserPageSetModel.
	UserPageSetModel interface {
		userPageSetModel
	}

	customUserPageSetModel struct {
		*defaultUserPageSetModel
	}
)

// NewUserPageSetModel returns a model for the database table.
func NewUserPageSetModel(conn sqlx.SqlConn, c cache.CacheConf) UserPageSetModel {
	return &customUserPageSetModel{
		defaultUserPageSetModel: newUserPageSetModel(conn, c),
	}
}
func (m *defaultUserPageSetModel) UpdateByUserID(ctx context.Context, newData *UserPageSet) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	chaosSystemUserPageSetIdKey := fmt.Sprintf("%s%v", cacheChaosSystemUserPageSetIdPrefix, data.Id)
	chaosSystemUserPageSetUserIdKey := fmt.Sprintf("%s%v", cacheChaosSystemUserPageSetUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userPageSetRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.UserId, newData.Avatar, newData.DefaultRouter, newData.SideMode, newData.TextColor, newData.ActiveTextColor, newData.Id)
	}, chaosSystemUserPageSetIdKey, chaosSystemUserPageSetUserIdKey)
	return err
}
