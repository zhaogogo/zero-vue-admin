package system

import (
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
