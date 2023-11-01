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
	//{"id":2,"host":"192.168.14.102","sliences":[{"id":0,"host_id":2,"slience_name":"a","default":true,"matchers":[{"name":"env","value":"aaa","is_regex":false,"is_equal":true,"host_id":2,"slience_name_id":0},{"name":"i","value":"aaa","is_regex":false,"is_equal":true,"host_id":2,"slience_name_id":0}]},{"id":0,"host_id":2,"slience_name":"b","default":false,"matchers":[{"name":"env","value":"bbb","is_regex":false,"is_equal":true,"host_id":2,"slience_name_id":0},{"name":"instance","value":"bbb","is_regex":false,"is_equal":true,"host_id":2,"slience_name_id":0}]}]}
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
