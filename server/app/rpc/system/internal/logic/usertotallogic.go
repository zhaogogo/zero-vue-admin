package logic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserTotalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserTotalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserTotalLogic {
	return &UserTotalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserTotalLogic) UserTotal(in *pb.Empty) (*pb.Total, error) {
	total, err := l.svcCtx.UserModel.Total_NC(l.ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "数据库查询失败")
	}

	return &pb.Total{Total: total}, nil
}
