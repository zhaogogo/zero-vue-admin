package hosts

import (
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/validate"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/logic/monitoring/hosts"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
)

func HandlerHostsSlienceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HandlerHostsSlienceRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		err := validate.StructExceptCtx(r.Context(), req)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		l := hosts.NewHandlerHostsSlienceLogic(r.Context(), svcCtx)
		resp, err := l.HandlerHostsSlience(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
