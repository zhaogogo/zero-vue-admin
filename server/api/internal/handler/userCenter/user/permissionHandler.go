package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/api/internal/logic/userCenter/user"
	"github.com/zhaoqiang0201/zero-vue-admin/server/api/internal/svc"
)

func PermissionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewPermissionLogic(r.Context(), svcCtx)
		resp, err := l.Permission()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
