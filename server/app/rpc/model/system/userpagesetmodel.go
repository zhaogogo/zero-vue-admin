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
		TransDeleteByUserID(ctx context.Context, session sqlx.Session, userid uint64) error
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

func (m *defaultUserPageSetModel) TransDeleteByUserID(ctx context.Context, session sqlx.Session, userid uint64) error {
	data, err := m.FindOneByUserId(ctx, userid)
	if err != nil {
		return err
	}

	chaosSystemUserPageSetIdKey := fmt.Sprintf("%s%v", cacheChaosSystemUserPageSetIdPrefix, data.Id)
	chaosSystemUserPageSetUserIdKey := fmt.Sprintf("%s%v", cacheChaosSystemUserPageSetUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
		return session.ExecCtx(ctx, query, userid)
	}, chaosSystemUserPageSetIdKey, chaosSystemUserPageSetUserIdKey)
	return err
}
