package logic

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/rpc/model/systemusermodel"
	"time"

	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/rpc/system/pb"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/rpc/system/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.UserName)
	if err != nil {
		if errors.Is(err, systemusermodel.ErrNotFound) {
			return nil, errors.Wrapf(err, "用户不存在, 参数: %s", in.UserName)
		}
		return nil, errors.Wrapf(err, "用户登陆失败, 参数: %s", in.UserName)
	}
	if err := utils.CheckPassword(in.PassWord, user.Password); err != nil {
		return nil, errors.Wrapf(err, "用户密码错误")
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	jwtToken, err := l.getToken(l.svcCtx.Config.JWT.AccessSecret, now, accessExpire, user.Id)
	if err != nil {
		return nil, errors.Wrapf(err, "生产token失败, 参数: %v", in)
	}
	resp := &pb.LoginResponse{
		ID:           user.Id,
		UserName:     user.Name,
		Token:        jwtToken,
		ExpireAt:     now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}

	return resp, nil
}

func (l *LoginLogic) getToken(secretKey string, iat, seconds, userid int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userid
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
