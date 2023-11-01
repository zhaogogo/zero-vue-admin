package hosts

import (
	"fmt"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/validate"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/logic/monitoring/hosts"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
)

func CreateSlienceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SliencePutRequest
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Println("xxx", err)
			httpx.Error(w, err)
			return
		}
		err := validate.StructExceptCtx(r.Context(), req)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		l := hosts.NewCreateSlienceLogic(r.Context(), svcCtx)
		resp, err := l.CreateSlience(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
