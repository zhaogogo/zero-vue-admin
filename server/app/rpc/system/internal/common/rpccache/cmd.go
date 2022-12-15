package rpccache

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/model/system"
)

func SetCacheWithNotFound(ctx context.Context, redis *redis.Redis, key string) error {
	return redis.SetexCtx(ctx, key, "*", int(system.unstable.AroundDuration(system.option.NotFoundExpiry).Seconds()))
}
