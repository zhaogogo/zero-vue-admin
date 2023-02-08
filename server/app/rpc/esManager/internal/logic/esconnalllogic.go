package logic

import (
	"context"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ESConnAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewESConnAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ESConnAllLogic {
	return &ESConnAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ESConnAllLogic) ESConnAll(in *pb.Empty) (*pb.ESConnResponse, error) {
	mesconns, err := l.svcCtx.ESConnModel.FindAll_NC(l.ctx)
	if err != nil {
		return nil, err
	}
	pesconns := []*pb.ESConn{}
	for _, esconn := range mesconns {
		p := &pb.ESConn{
			ID:       esconn.Id,
			ESConn:   esconn.EsConn,
			Version:  esconn.Version,
			User:     esconn.User.String,
			PassWord: esconn.Password.String,
			Describe: esconn.Describe,
		}

		pesconns = append(pesconns, p)
	}

	return &pb.ESConnResponse{
		ESConns: pesconns,
	}, nil
}
