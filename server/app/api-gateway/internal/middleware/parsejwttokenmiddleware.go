package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"net/http"
	"strconv"
)

type ParseJWTTokenMiddleware struct {
}

func NewParseJWTTokenMiddleware() *ParseJWTTokenMiddleware {
	return &ParseJWTTokenMiddleware{}
}

func (m *ParseJWTTokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rCtx := r.Context()
		userid_i := rCtx.Value("userID")
		if userid_i == nil {
			httpx.Error(
				w,
				errorx.NewByCode(errors.New(
					errorx.ErrorxMessage(errorx.JWT_TOKEN_PARAM_GET_NERROR)),
					errorx.JWT_TOKEN_PARAM_GET_NERROR).WithMeta("", "", map[string]string{"jwt_claims": "jwt token无userID claims"}),
			)
			return
		}
		userid_jsonNumber, ok := userid_i.(json.Number)
		if !ok {
			httpx.Error(
				w,
				errorx.NewByCode(errors.New(
					errorx.ErrorxMessage(errorx.JWT_TOKEN_PARAM_GET_NERROR)),
					errorx.JWT_TOKEN_PARAM_GET_NERROR).WithMeta("", "", map[string]string{"jwt_claims": "userID", "jwt断言": fmt.Sprintf("want: %T, got: json.Number", userid_i)}),
			)
			return
		}
		userid, err := strconv.ParseUint(userid_jsonNumber.String(), 10, 64)
		if err != nil {
			httpx.Error(
				w,
				errorx.NewByCode(errors.New(
					errorx.ErrorxMessage(errorx.JWT_TOKEN_PARAM_GET_NERROR)),
					errorx.JWT_TOKEN_PARAM_GET_NERROR).WithMeta("", "", map[string]string{"jwt_claims": "userID", "jwt类型转换错误": fmt.Sprintf(" value: %v, error: %v", userid_jsonNumber.String(), err)}),
			)
			return
		}
		// ------------
		username_i := rCtx.Value("userName")
		if username_i == nil {
			httpx.Error(
				w,
				errorx.NewByCode(errors.New(
					errorx.ErrorxMessage(errorx.JWT_TOKEN_PARAM_GET_NERROR)),
					errorx.JWT_TOKEN_PARAM_GET_NERROR).WithMeta("", "", map[string]string{"jwt_claims": "jwt token无userName claims"}),
			)
			return
		}
		if !ok {
			httpx.Error(
				w,
				errorx.NewByCode(errors.New(
					errorx.ErrorxMessage(errorx.JWT_TOKEN_PARAM_GET_NERROR)),
					errorx.JWT_TOKEN_PARAM_GET_NERROR).WithMeta("", "", map[string]string{"jwt_claims": "userName", "jwt断言": fmt.Sprintf("want: %T, got: string", username_i)}),
			)
			return
		}

		ctx := context.WithValue(rCtx, "user_id", userid)
		next(w, r.WithContext(ctx))
	}
}
