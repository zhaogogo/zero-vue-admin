package logic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ESConnDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewESConnDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ESConnDetailLogic {
	return &ESConnDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ESConnDetailLogic) ESConnDetail(in *pb.ESConnID) (*pb.ESConn, error) {
	mesconn, err := l.svcCtx.ESConnModel.FindOne(l.ctx, in.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "数据库查询错误")
	}

	return &pb.ESConn{
		ID:       mesconn.Id,
		ESConn:   mesconn.EsConn,
		Version:  mesconn.Version,
		User:     mesconn.User.String,
		PassWord: mesconn.Password.String,
		Describe: mesconn.Describe,
	}, nil
}
