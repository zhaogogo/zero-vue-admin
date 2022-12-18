package middleware

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/casbinx"
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
		var (
			hasPerimisstion bool
			err             error
		)
		userid := r.Context().Value("user_id").(uint64)
		params := &systemservice.UserID{ID: userid}
		userRoleList, err := middleSvcCtx.Ctx.SystemRpcClient.UserRoleByUserID(r.Context(), params)
		if err != nil {
			s, _ := status.FromError(err)
			if s.Message() == sql.ErrNoRows.Error() {
				userRoleList = new(systemservice.UserRoleResponse)
				goto TOPERMISSION
			}
			httpx.Error(w, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("*SystemRpcClient.UserRoleByUserID", err.Error(), params))
			return
		}
	TOPERMISSION:
		// 获取请求的URI
		obj := r.URL.RequestURI() // r.RequestURI    /api/v1/system/usermenus?k1=v1&k2=v2
		// 获取请求方法
		act := r.Method
		// 获取用户的角色
		for _, userRole := range userRoleList.UserRoles {
			hasPerimisstion, err = casbinx.Casbin.Enforce(strconv.FormatUint(userRole.RoleID, 10), obj, act)
			if err != nil {
				logx.Errorf("roleid: %v, error: %v", r, err)
			}
			fmt.Println("=====>", userRole.RoleID, obj, act, hasPerimisstion)
			if hasPerimisstion {
				break
			}
		}
		//没有权限
		if !hasPerimisstion {
			httpx.Error(w, errorx.NewByCode(errors.New("用户权限被拒绝"), errorx.USER_PERMISSION_REJECT))
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, "c-UserID", userid)
		next(w, r.WithContext(ctx))
	}
}
