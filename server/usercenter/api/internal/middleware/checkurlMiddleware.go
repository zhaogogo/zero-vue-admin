package middleware

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/api/pkg/errorx"
	"net/http"
)

type CheckUrlMiddleware struct {
	Redis *redis.Redis
}

func NewCheckUrlMiddleware(redis *redis.Redis) *CheckUrlMiddleware {
	return &CheckUrlMiddleware{Redis: redis}
}

func (m *CheckUrlMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 获取jwt token的userId
		// github.com/zeromicro/go-zero@v1.4.0/rest/handler/authhandler.go:77
		userId := r.Context().Value("userId").(json.Number).String()
		if userId == "" {
			httpx.Error(w, errorx.New(errorx.SERVER_COMMON_ERROR, "jwt token无userId"))
			return
		}
		next(w, r)
	}
}
