package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserDetailLogic) UserDetail(in *pb.UserID) (*pb.User, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.ID)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, err
		}
		return nil, errors.Wrap(err, "数据库查询失败")
	}
	res := &pb.User{
		ID:         user.Id,
		Name:       user.Name,
		NickName:   user.NickName,
		PassWord:   user.Password,
		UserType:   user.Type,
		Email:      user.Email,
		Phone:      user.Phone,
		Department: user.Department,
		Position:   user.Position,
		CreateBy:   user.CreateBy,
		CreateTime: user.CreateTime.Unix(),
		UpdateBy:   user.UpdateBy,
		UpdateTime: user.UpdateTime.Unix(),
		DeleteBy:   user.DeleteBy,
	}
	if user.DeleteTime.Valid {
		res.DeleteTime = user.DeleteTime.Time.Unix()
	} else {
		res.DeleteTime = 0
	}

	return res, nil
}
