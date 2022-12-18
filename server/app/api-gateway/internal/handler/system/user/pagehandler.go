package user

import (
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/validate"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/logic/system/user"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
)

func PageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserPageRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		if err := validate.StructExceptCtx(r.Context(), req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := user.NewPageLogic(r.Context(), svcCtx)
		resp, err := l.Page(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
