package hosts

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"gorm.io/gorm/clause"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSlienceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSlienceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSlienceLogic {
	return &CreateSlienceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSlienceLogic) CreateSlience(req *types.SliencePutRequest) (resp *types.HttpCommonResponse, err error) {
	err = l.svcCtx.MonitoringDB().Clauses(clause.OnConflict{UpdateAll: true}).Create(&req.Sliences).Error
	if err != nil {
		return nil, errorx.New(err, "创建静默规则失败")
	}
	return &types.HttpCommonResponse{
		Code: 200,
		Msg:  "ok",
		Meta: nil,
	}, nil
}
