package logic

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/common/utils"
	"time"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

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
	user, err := l.svcCtx.UserModel.FindOneByNameWHEREDeleteTimeISNULL(l.ctx, in.Name)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, errors.Wrapf(err, "无数据, 表: user, 字段: name=%s", in.Name)
		}
		return nil, errors.Wrapf(err, "查询用户失败, 表: user, 字段: name=%s", in.Name)
	}
	if err := utils.CheckPassword(in.PassWord, user.Password); err != nil {
		return nil, errors.Wrapf(err, "密码错误, 用户: %s", in.Name)
	}
	roleid := []uint64{}
	userroles, _ := l.svcCtx.UserRoleModel.FindByUserID(l.ctx, l.svcCtx.Redis, user.Id)
	for _, userrole := range userroles {
		roleid = append(roleid, userrole.RoleId)
	}

	token, expire, refresh, err := l.genJWTToken(user.Id, roleid)
	if err != nil {
		return nil, errors.Wrapf(err, "生成token失败, 用户: %s", in.Name)
	}
	return &pb.LoginResponse{
		Token:        token,
		ExporeAt:     expire,
		RefreshAfter: refresh,
		UserId:       user.Id,
	}, nil
}

/*
iss (issuer)：签发人
exp (expiration time)：过期时间
sub (subject)：主题
aud (audience)：受众
nbf (Not Before)：生效时间
iat (Issued At)：签发时间
jti (JWT ID)：编号
*/
func (l *LoginLogic) genJWTToken(userID uint64, roleid []uint64) (t string, expire int64, refresh int64, err error) {
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	claims := make(jwt.MapClaims)
	claims["exp"] = now + accessExpire
	claims["iat"] = now
	claims["userID"] = userID
	claims["roleIDs"] = roleid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	t, err = token.SignedString([]byte(l.svcCtx.Config.JWT.AccessSecret))
	if err != nil {
		return "", 0, 0, err
	}
	return t, now + accessExpire, now + accessExpire/2, nil
}
