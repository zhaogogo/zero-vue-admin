package connect

import (
	"context"
	"fmt"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/model/monitoring"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.ConnectManagerCreateRequest) (resp *types.HttpCommonResponse, err error) {
	fmt.Printf("%#v", req)
	conn := monitoring.StoreConnectManager{
		Type:      req.Type,
		Env:       req.Env,
		Host:      req.Host,
		AccessKey: req.AccessKey,
		SecretKey: req.SecretKey,
	}

	if err := l.svcCtx.MonitoringDB().Create(&conn).Error; err != nil {
		return nil, errorx.New(err, "创建连接失败")
	}

	return &types.HttpCommonResponse{Code: 200, Msg: "OK", Meta: nil}, nil
}
