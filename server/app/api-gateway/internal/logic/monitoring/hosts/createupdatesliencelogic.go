package hosts

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/slience"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpdateSlienceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUpdateSlienceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpdateSlienceLogic {
	return &CreateUpdateSlienceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUpdateSlienceLogic) CreateUpdateSlience(req *types.SliencePutRequest) (resp *types.HttpCommonResponse, err error) {
	err = l.svcCtx.MonitoringDB().Transaction(func(tx *gorm.DB) error {
		// 替换时SlienceNames，做替换和删除的动作 SlienceName下的Matchers只做替换 不删除 ，需要下面的for循环做删除
		logx.Severef("xxx-%d", 1)
		err := tx.Model(&types.Host{Id: req.ID}).Session(&gorm.Session{FullSaveAssociations: true}).Select("").Association("SlienceNames").Unscoped().Replace(req.Sliences)
		if err != nil {
			return errors.Wrapf(err, "关联查询Host替换SlienceNames失败, where (id: %v, host: %s), replace: (%#v)", req.ID, req.Host, req.Sliences)
		}
		for _, slienceName := range req.Sliences {
			s := slienceName
			err = tx.Model(&types.SlienceName{Id: s.Id}).Session(&gorm.Session{FullSaveAssociations: true}).Association("Matchers").Unscoped().Replace(s.Matchers)
			if err != nil {
				return errors.Wrapf(err, "关联查询SlienceNames替换Matchers失败,where (%#v)", s)
			}
		}
		return nil
	})
	if err != nil {
		return nil, errorx.New(err, "关联查询替换失败")
	}
	err = slience.GetConsumerSliences(l.svcCtx.MonitoringDB(), l.svcCtx.SlienceList)
	if err != nil {
		logx.Errorf("更新Slience，刷新全局规则失败， error: %v", err)
	}
	return &types.HttpCommonResponse{
		Code: 200,
		Msg:  "ok",
		Meta: nil,
	}, nil
}
