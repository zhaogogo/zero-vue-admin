package pageset

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/logic/usercenter/user/pageset"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/svc"
)

func GetUserPageSetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := pageset.NewGetUserPageSetLogic(r.Context(), svcCtx)
		resp, err := l.GetUserPageSet()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
