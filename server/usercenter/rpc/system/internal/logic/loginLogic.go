package logic

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/rpc/system/pkg/utils"
	"time"

	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/rpc/system/pb"

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
	userinfo, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.UserName)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, errors.Wrap(err, "用户不存在")
		}
		return nil, errors.Wrap(err, "数据库查询错误")
	}
	if err := utils.CheckPassword(in.PassWord, userinfo.Password); err != nil {
		return nil, errors.Wrap(err, "密码错误")
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	jwttoken, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, userinfo.Id)
	if err != nil {
		return nil, errors.Wrap(err, "生成token失败")
	}
	res := &pb.LoginResponse{
		Token:        jwttoken,
		ExpireAt:     accessExpire,
		RefreshAfter: now + accessExpire/2,
	}

	return res, nil
}
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
