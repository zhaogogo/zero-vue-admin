package logic

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/model"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *pb.UpdateUserRequest) (*pb.Empty, error) {
	err := l.svcCtx.UserModel.UpdateWithOutPassword(l.ctx, &model.User{
		Id:       in.ID,
		Name:     in.Name,
		NickName: in.NickName,
		//Password   string        `db:"password"`    // 密码
		//Type       int64         `db:"type"`        // 账户类型 0-本地用户 1-ldap用户
		Email:      in.Email,
		Phone:      in.Phone,
		Department: in.Department,
		Position:   in.Position,
		//CreateBy   string        `db:"create_by"`   // 创建人
		//CreateTime time.Time     `db:"create_time"` // 创建时间
		UpdateBy: in.UpdateBy,
		//UpdateTime time.Time     `db:"update_time"` // 更新时间
		//DeleteBy   string        `db:"delete_by"`   // 删除人
		//DeleteTime sql.NullTime  `db:"delete_time"` // 删除时间
		//PageSetId  sql.NullInt64 `db:"page_set_id"`
	})
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
