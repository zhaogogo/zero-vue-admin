package middleware

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/middleSvcCtx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"net/http"
)

type CheckUserExistsMiddleware struct {
}

func NewCheckUserExistsMiddleware() *CheckUserExistsMiddleware {
	return &CheckUserExistsMiddleware{}
}

func (m *CheckUserExistsMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userid, err := r.Context().Value("userID").(json.Number).Int64()
		if err != nil {
			w.Header().Set(httpx.ContentType, "application/json; charset=utf-8")
			v := responseerror.ErrorResponse{
				Code:     int32(errorx.UNAUTHORIZATION),
				Msg:      errorx.ErrorxMessage(errorx.UNAUTHORIZATION),
				CauseErr: errorx.ErrorxMessage(errorx.UNAUTHORIZATION),
			}
			bs, err := json.Marshal(v)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(401)
			w.Write(bs)
			return
		}

		userinfo, err := middleSvcCtx.Ctx.SystemRpcClient.UserInfo(r.Context(), &systemservice.UserID{ID: uint64(userid)})
		if err != nil || userinfo.DeleteTime != 0 {
			w.Header().Set(httpx.ContentType, "application/json; charset=utf-8")
			v := responseerror.ErrorResponse{
				Code:     int32(errorx.UNAUTHORIZATION),
				Msg:      errorx.ErrorxMessage(errorx.UNAUTHORIZATION),
				CauseErr: errorx.ErrorxMessage(errorx.UNAUTHORIZATION),
			}
			bs, err := json.Marshal(v)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(401)
			w.Write(bs)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, "c-UserName", userinfo.Name)
		next(w, r.WithContext(ctx))
	}
}
