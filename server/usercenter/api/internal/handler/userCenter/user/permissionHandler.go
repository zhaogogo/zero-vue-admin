package user

import (
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/api/internal/logic/userCenter/user"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/api/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
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
