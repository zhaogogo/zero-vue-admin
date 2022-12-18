package middleware

import (
	"database/sql"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/middleSvcCtx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"google.golang.org/grpc/status"
	"net/http"
)

type CheckUserExistsMiddleware struct {
}

func NewCheckUserExistsMiddleware() *CheckUserExistsMiddleware {
	return &CheckUserExistsMiddleware{}
}

func (m *CheckUserExistsMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userid := r.Context().Value("user_id").(uint64)
		userinfoParam := &systemservice.UserID{ID: uint64(userid)}
		userinfo, err := middleSvcCtx.Ctx.SystemRpcClient.UserDetail(r.Context(), userinfoParam)
		if err != nil {
			s, _ := status.FromError(err)
			if s.Message() == sql.ErrNoRows.Error() {
				httpx.Error(w, errorx.NewByCode(err, errorx.UNAUTHORIZATION).WithMeta("SystemRpcClient.UserInfo", err.Error(), userinfoParam))
				return
			}
			httpx.Error(w, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("SystemRpcClient.UserInfo", err.Error(), userinfoParam))
			return
		}
		if userinfo.DeleteTime != 0 {
			httpx.Error(w, errorx.NewByCode(errors.New("用户为禁用状态"), errorx.UNAUTHORIZATION))
			return
		}

		next(w, r)
	}
}
