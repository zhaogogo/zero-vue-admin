package pageset

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/logic/usercenter/user/pageset"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/types"
)

func SetUserPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserPageRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := pageset.NewSetUserPageLogic(r.Context(), svcCtx)
		resp, err := l.SetUserPage(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
