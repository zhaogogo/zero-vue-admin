package logic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/common/utils"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/model"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserUserRoleLogic {
	return &CreateUserUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserUserRoleLogic) CreateUser_UserRole(in *pb.CreateUser_UserRoleRequest) (*pb.Empty, error) {
	pass, err := utils.GenPassword(in.User.PassWord)
	if err != nil {
		return nil, errors.Wrap(err, "生成密码失败")
	}
	err = l.svcCtx.UserModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		user := &model.User{
			Id:         0,
			Name:       in.User.Name,
			NickName:   in.User.NickName,
			Password:   pass,
			Type:       0,
			Email:      in.User.Email,
			Phone:      in.User.Phone,
			Department: in.User.Department,
			Position:   in.User.Position,
			CreateBy:   in.User.CreateBy,
			UpdateBy:   in.User.CreateBy,
			DeleteBy:   "",
			DeleteTime: sql.NullTime{},
			PageSetId:  sql.NullInt64{},
		}
		r, err := l.svcCtx.UserModel.TransInsert(ctx, session, user)
		if err != nil {
			return errors.Wrap(err, "插入用户失败")
		}
		userid, err := r.LastInsertId()
		if err != nil {
			return errors.Wrap(err, "获取插入用户id失败")
		}

		err = l.svcCtx.UserRoleModel.TransDeleteByUserID(ctx, session, uint64(userid))
		if err != nil {
			return errors.Wrap(err, "删除用户角色失败")
		}

		if len(in.RoleList) == 0 {
			return nil
		}

		r, err = l.svcCtx.UserRoleModel.TranInsertUserIDRoleIDs(ctx, session, uint64(userid), in.RoleList)
		if err != nil {
			return err
		}
		a, err := r.RowsAffected()
		if err != nil {
			return errors.Wrap(err, "获取影响行数错误")
		}
		if int64(len(in.RoleList)) != a {
			return errors.Wrap(err, "插入行数错误")
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
