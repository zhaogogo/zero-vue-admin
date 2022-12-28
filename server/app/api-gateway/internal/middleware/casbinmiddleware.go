package middleware

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/middleSvcCtx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
)

type CasbinMiddleware struct {
}

func NewCasbinMiddleware() *CasbinMiddleware {
	return &CasbinMiddleware{}
}

func (m *CasbinMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userid := r.Context().Value("user_id").(uint64)
		userinfo := r.Context().Value("userinfo").(*systemservice.User)
		params := &systemservice.UserID{ID: userid}
		userRoleList, err := middleSvcCtx.Ctx.SystemRpcClient.UserRoleByUserID(r.Context(), params)
		if err != nil {
			s, _ := status.FromError(err)
			if s.Message() == sql.ErrNoRows.Error() {
				//httpx.Error(w, errorx.NewByCode(errors.New("用户权限被拒绝"), errorx.USER_PERMISSION_REJECT))
				userRoleList = new(systemservice.UserRoleResponse)
				goto PERMISSION
			}
			httpx.Error(w, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("*SystemRpcClient.UserRoleByUserID", err.Error(), params))
			return
		}
	PERMISSION:
		ctx := r.Context()
		ctx = context.WithValue(ctx, "userroleinfo", userRoleList.UserRoles)
		// 获取请求的URI
		obj := r.URL.RequestURI() // r.RequestURI    /api/v1/system/usermenus?k1=v1&k2=v2
		// 获取请求方法
		act := r.Method
		// 获取用户的角色

		casbinEnforcerParam := &systemservice.CasbinEnforceRequest{Sub: strconv.FormatUint(userinfo.CurrentRole, 10), Obj: obj, Act: act}
		res, err := middleSvcCtx.Ctx.SystemRpcClient.CasbinEnforcer(r.Context(), casbinEnforcerParam)
		if err != nil {
			httpx.Error(w, errorx.New(err, "CasbinEnforcer错误").WithMeta("SystemRpcClient.CasbinEnforcer", err.Error(), casbinEnforcerParam))
			return
		}
		fmt.Println("=====>", userinfo.CurrentRole, obj, act, res.Success)

		//没有权限
		if !res.Success {
			httpx.Error(w, errorx.NewByCode(errors.New("用户权限被拒绝"), errorx.USER_PERMISSION_REJECT))
			return
		}
		next(w, r.WithContext(ctx))
	}
}
