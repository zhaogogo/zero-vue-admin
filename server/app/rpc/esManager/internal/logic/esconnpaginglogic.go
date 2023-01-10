package logic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ESConnPagingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewESConnPagingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ESConnPagingLogic {
	return &ESConnPagingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ESConnPagingLogic) ESConnPaging(in *pb.ESConnPagingRequest) (*pb.ESConnResponse, error) {
	res, err := l.svcCtx.ESConnModel.FindPaging_NC(l.ctx, in.Page, in.PageSize)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "数据库查询错误")
	}
	pesConn := []*pb.ESConn{}
	for _, v := range res {
		pesConn = append(pesConn, &pb.ESConn{
			ID:       v.Id,
			ESConn:   v.EsConn,
			Version:  v.Version,
			User:     v.User.String,
			PassWord: v.User.String,
			Describe: v.Describe,
		})
	}
	return &pb.ESConnResponse{
		ESConns: pesConn,
	}, nil
}
