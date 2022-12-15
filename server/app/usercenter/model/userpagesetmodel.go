package model

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
		CreateODuplicateByUserId(ctx context.Context, newData *UserPageSet, userID uint64) error
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

func (m *defaultUserPageSetModel) CreateODuplicateByUserId(ctx context.Context, newData *UserPageSet, userID uint64) error {
	usercenterUserPageSetUserIdKey := fmt.Sprintf("%s%v", cacheUsercenterUserPageSetUserIdPrefix, userID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("INSERT INTO %s(%s) VALUES(?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE `user_id`=VALUES(`user_id`), `avatar`=VALUES(`avatar`), `default_router`=VALUES(`default_router`), `side_mode`=VALUES(`side_mode`), `text_color`=VALUES(`text_color`), `active_text_color`=VALUES(`active_text_color`)", m.table, userPageSetRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, userID, newData.Avatar, newData.DefaultRouter, newData.SideMode, newData.TextColor, newData.ActiveTextColor)
	}, usercenterUserPageSetUserIdKey)
	return err
}
