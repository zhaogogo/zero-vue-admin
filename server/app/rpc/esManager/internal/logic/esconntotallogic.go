package logic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ESConnTotalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewESConnTotalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ESConnTotalLogic {
	return &ESConnTotalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ESConnTotalLogic) ESConnTotal(in *pb.Empty) (*pb.Total, error) {
	total, err := l.svcCtx.ESConnModel.FindTotal_NC(l.ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "数据库查询错误")
	}

	return &pb.Total{
		Total: total,
	}, nil
}
