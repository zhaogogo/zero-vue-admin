package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/logic/system/user"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
)

func AllUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewAllUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.AllUserInfo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
