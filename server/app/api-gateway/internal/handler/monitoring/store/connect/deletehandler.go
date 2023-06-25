package connect

import (
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/logic/monitoring/store/connect"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/validate"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
)

func DeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ConnectManagerDeleteRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		if err := validate.StructExceptCtx(r.Context(), req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := connect.NewDeleteLogic(r.Context(), svcCtx)
		resp, err := l.Delete(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
