package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	cacheUsercenterUserRoleUserIdPrefix = "cache:usercenter:userRole:user_id:"
)

var _ UserRoleModel = (*customUserRoleModel)(nil)

type (
	// UserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserRoleModel.
	UserRoleModel interface {
		userRoleModel
		FindByUserID(context.Context, uint64) ([]uint64, error)
	}

	customUserRoleModel struct {
		*defaultUserRoleModel
	}
)

// NewUserRoleModel returns a model for the database table.
func NewUserRoleModel(conn sqlx.SqlConn, c cache.CacheConf) UserRoleModel {
	return &customUserRoleModel{
		defaultUserRoleModel: newUserRoleModel(conn, c),
	}
}

func (m *defaultUserRoleModel) FindByUserID(ctx context.Context, userID uint64) ([]uint64, error) {
	usercenterUserRoleUserIDKey := fmt.Sprintf("%s%v", cacheUsercenterUserRoleUserIdPrefix, userID)
	var (
		resp    []UserRole
		roleIDs []uint64
	)
	err := m.QueryRowCtx(ctx, &resp, usercenterUserRoleUserIDKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("SELECT %s FROM %s WHERE `user_id` = ?", userRoleRows, m.table)
		return conn.QueryRowsCtx(ctx, v, query, userID)
	})
	switch err {
	case nil:
		for _, r := range resp {
			roleIDs = append(roleIDs, r.RoleId)
		}
		return roleIDs, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
