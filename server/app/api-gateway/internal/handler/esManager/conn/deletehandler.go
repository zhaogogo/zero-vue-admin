package conn

import (
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/validate"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/logic/esManager/conn"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
)

func DELETEHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ESConnDeleteRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		if err := validate.StructExceptCtx(r.Context(), req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := conn.NewDELETELogic(r.Context(), svcCtx)
		resp, err := l.DELETE(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
