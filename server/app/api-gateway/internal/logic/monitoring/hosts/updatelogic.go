package hosts

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/slience"
	"gorm.io/gorm"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.Host) (resp *types.HttpCommonResponse, err error) {
	err = l.svcCtx.MonitoringDB().Transaction(func(tx *gorm.DB) error {
		err = tx.Session(&gorm.Session{FullSaveAssociations: true}).Updates(req).Error
		if err != nil {
			return errors.Wrap(err, "更新host失败")
		}
		err = slience.GetConsumerSliences(l.svcCtx.MonitoringDB(), l.svcCtx.SlienceList)
		if err != nil {
			return errors.Wrap(err, "刷新全局规则失败")
		}
		return nil
	})
	if err != nil {
		return nil, errorx.New(err, fmt.Sprintf("更新id: %d 失败", req.Id))
	}
	return &types.HttpCommonResponse{
		Code: 200,
		Msg:  "OK",
		Meta: nil,
	}, nil
}
