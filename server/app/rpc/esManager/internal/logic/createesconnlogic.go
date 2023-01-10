package logic

import (
	"context"
	"database/sql"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/model"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateESConnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateESConnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateESConnLogic {
	return &CreateESConnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateESConnLogic) CreateESConn(in *pb.CreateESConnRequest) (*pb.Empty, error) {
	_, err := l.svcCtx.ESConnModel.Insert(l.ctx, &model.EsConn{
		Id:       0,
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
