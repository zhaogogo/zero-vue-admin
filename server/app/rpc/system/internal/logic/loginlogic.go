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
			return nil, err
		}
		return nil, errors.Wrap(err, "数据库错误")
	}
	if err := utils.CheckPassword(in.PassWord, user.Password); err != nil {
		return nil, errors.Wrap(err, "密码错误")
	}
	//roleid := []uint64{}
	//userroles, err := l.svcCtx.UserRoleModel.FindByUserID(l.ctx, l.svcCtx.Redis, user.Id)
	//if err != nil {
	//	logx.Error(err)
	//} else {
	//	for _, userrole := range userroles {
	//		roleid = append(roleid, userrole.RoleId)
	//	}
	//}

	token, expire, refresh, err := l.genJWTToken(user.Id, user.Name)
	if err != nil {
		return nil, errors.Wrap(err, "生成token失败")
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
func (l *LoginLogic) genJWTToken(userID uint64, userName string) (t string, expire int64, refresh int64, err error) {
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	claims := make(jwt.MapClaims)
	claims["exp"] = now + accessExpire
	claims["iat"] = now
	claims["userID"] = userID
	claims["userName"] = userName
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	t, err = token.SignedString([]byte(l.svcCtx.Config.JWT.AccessSecret))
	if err != nil {
		return "", 0, 0, err
	}
	return t, now + accessExpire, now + accessExpire/2, nil
}
