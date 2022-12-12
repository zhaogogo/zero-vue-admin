package user

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	errorx2 "github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/common/utils"
	"time"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	user, err := l.svcCtx.UserModel.FindOneByNameWHEREDeleteTimeISNULL(l.ctx, req.UserName)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, errorx2.NewFromCode(errorx2.USER_NOT_FOUND, err)
		}
		return nil, errorx2.NewFromCode(errorx2.DB_ERROR, err)
	}
	if err := utils.CheckPassword(req.PassWord, user.Password); err != nil {
		logx.Errorf("密码错误, error: %v", err)
		return nil, errorx2.New(errorx2.SERVER_COMMON_ERROR, "用户密码错误")
	}
	token, expire, refresh, err := l.genJWTToken(user.Id)
	if err != nil {
		logx.Errorf("生成jwt token失败, error: %v", err)
		return nil, errorx2.New(errorx2.SERVER_COMMON_ERROR, "生成jwt token失败")
	}
	return &types.LoginResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK"},
		Token:              token,
		ExpireAt:           expire,
		RefreshAfter:       refresh,
	}, nil
}

func (l *LoginLogic) genJWTToken(userID uint64) (t string, expire int64, refresh int64, err error) {
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	claims := make(jwt.MapClaims)
	claims["exp"] = now + accessExpire
	claims["iat"] = now
	claims["userID"] = userID
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	t, err = token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
	if err != nil {
		return "", 0, 0, err
	}
	return t, accessExpire, now + accessExpire/2, nil
}
