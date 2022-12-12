package middleware

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/casbinx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror"
	"net/http"
)

type CasbinMiddleware struct {
}

func NewCasbinMiddleware() *CasbinMiddleware {
	return &CasbinMiddleware{}
}

func (m *CasbinMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//userid, err := r.Context().Value("userID").(json.Number).Int64()
		roleids, ok := r.Context().Value("roleIDs").([]interface{})
		if !ok {
			logx.Error("解析JWT携带参数roleIDs, 断言失败")
		}
		var (
			hasPerimisstion bool
			err             error
		)
		// 获取请求的URI
		obj := r.URL.RequestURI() // r.RequestURI    /api/v1/system/usermenus?k1=v1&k2=v2
		// 获取请求方法
		act := r.Method
		// 获取用户的角色

		for _, roleid := range roleids {
			r := roleid.(json.Number).String()
			hasPerimisstion, err = casbinx.Casbin.Enforce(r, obj, act)
			if err != nil {
				logx.Errorf("casbin权限验证失败, error: %v", err)
			}
			if hasPerimisstion {
				break
			}
		}
		//没有权限
		if !hasPerimisstion {
			w.Header().Set(httpx.ContentType, "application/json; charset=utf-8")
			v := responseerror.ErrorResponse{
				Code:     int32(errorx.USER_PERMISSION_REJECT),
				Msg:      errorx.ErrorxMsg(errorx.USER_PERMISSION_REJECT),
				CauseErr: errorx.ErrorxMsg(errorx.USER_PERMISSION_REJECT),
			}
			bs, err := json.Marshal(v)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(bs)
			return
		}
		next(w, r)
	}
}
