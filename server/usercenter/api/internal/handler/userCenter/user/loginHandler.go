package user

import (
	"fmt"
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
			//response.CustomHttpResponse(w, r, nil, err)
			httpx.Error(w, err)
			return
		}
		aaa := struct {
			Authorization  string   `json:"Authorization"`
			ContentType    string   `json:"Content-Type"`
			AcceptEncoding []string `json:"Accept-Encoding"`
			AAA            string   `json:"aaa"`
		}{}

		err := httpx.ParseHeaders(r, &aaa)
		fmt.Println("ParseHeaders error ==> ", err)
		fmt.Printf("ParseHeaders ==> %#v\n", aaa)

		//err := svcCtx.Validator.StructCtx(r.Context(), req);
		err = validate.StructExceptCtx(r.Context(), req)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		//response.CustomHttpResponse(w, r, resp, err)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
