package user

import (
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/api/internal/logic/userCenter/user"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/api/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/api/internal/types"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/pkg/validate"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		//err := svcCtx.Validator.StructCtx(r.Context(), req);
		err := validate.StructExceptCtx(r.Context(), req)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
