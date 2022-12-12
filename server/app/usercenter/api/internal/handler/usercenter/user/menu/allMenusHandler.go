package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/logic/usercenter/user/menu"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/svc"
)

func AllMenusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewAllMenusLogic(r.Context(), svcCtx)
		resp, err := l.AllMenus()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
