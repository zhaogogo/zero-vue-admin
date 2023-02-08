package es

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/pkg/es/es7"
)

type ES interface {
	Ping(ctx context.Context) (pingres interface{}, err error)
	Cat() CatInterface
}

func NewES(url string, user string, password string, version int64, id uint64) ES {
	switch version {
	case 7:
		return es7.NewES7(url, user, password, id)
	default:
		return nil
	}
}
