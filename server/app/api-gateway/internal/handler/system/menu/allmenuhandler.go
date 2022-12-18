package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/logic/system/menu"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
)

func AllMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewAllMenuLogic(r.Context(), svcCtx)
		resp, err := l.AllMenu()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
