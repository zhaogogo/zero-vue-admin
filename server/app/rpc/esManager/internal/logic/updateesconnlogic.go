package logic

import (
	"context"
	"database/sql"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/model"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateESConnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateESConnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateESConnLogic {
	return &UpdateESConnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateESConnLogic) UpdateESConn(in *pb.UpdateESConnRequest) (*pb.Empty, error) {
	err := l.svcCtx.ESConnModel.Update(l.ctx, &model.EsConn{
		Id:       in.ID,
		EsConn:   in.ESConn,
		Version:  in.Version,
		User:     sql.NullString{String: in.User, Valid: true},
		Password: sql.NullString{String: in.PassWord, Valid: true},
		Describe: in.Describe,
	})
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
