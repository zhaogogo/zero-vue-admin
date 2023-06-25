package file

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/logic/monitoring/store/file"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
)

func EnvSelectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := file.NewEnvSelectLogic(r.Context(), svcCtx)
		resp, err := l.EnvSelect()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
