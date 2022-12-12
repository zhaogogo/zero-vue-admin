package system

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
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
		FindByUserID(ctx context.Context, redis *redis.Redis, userID uint64) ([]UserRole, error)
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

func (m *defaultUserRoleModel) FindByUserID(ctx context.Context, redis *redis.Redis, userID uint64) ([]UserRole, error) {

	var resp []UserRole
	//-------------
	// redis查询
	//-------------
	usercenterUserRoleUserIDKey := fmt.Sprintf("%s%v", cacheUsercenterUserRoleUserIdPrefix, userID)
	//
	err := m.GetCacheCtx(ctx, usercenterUserRoleUserIDKey, &resp)
	fmt.Printf(">>>>>>>>> %T, %v\n", err, err)

	query := fmt.Sprintf("SELECT %s FROM %s WHERE `user_id` = ?", userRoleRows, m.table)
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, userID)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, ErrNotFound
	}
	m.SetCacheCtx(ctx, usercenterUserRoleUserIDKey, resp)
	return resp, nil

}
