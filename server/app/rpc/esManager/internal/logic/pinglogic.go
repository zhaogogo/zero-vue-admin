package logic

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/pkg/es"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *pb.PingRequest) (*pb.PingResponse, error) {
	esConn, err := l.svcCtx.ESConnModel.FindOne(l.ctx, in.EsConnID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "数据库错误")
	}
	e := es.NewES(esConn.EsConn, esConn.User.String, esConn.Password.String, esConn.Version, esConn.Id)
	if e == nil {
		return nil, errors.New("只支持7.X版本")
	}
	pingRes, err := e.Ping(l.ctx)
	if err != nil {
		return nil, err
	}
	r, err := json.Marshal(pingRes)
	if err != nil {
		return nil, err
	}

	return &pb.PingResponse{
		Data: &anypb.Any{Value: r},
	}, nil
}
