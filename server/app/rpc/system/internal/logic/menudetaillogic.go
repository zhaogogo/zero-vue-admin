package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuDetailLogic {
	return &MenuDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuDetailLogic) MenuDetail(in *pb.MenuID) (*pb.Menu, error) {
	menu, err := l.svcCtx.MenuModel.FindOne(l.ctx, in.ID)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, err
		}
		return nil, errors.Wrap(err, "数据库查询失败")
	}
	pbMenu := &pb.Menu{
		ID:         menu.Id,
		ParentID:   menu.ParentId,
		Name:       menu.Name,
		Path:       menu.Path,
		Component:  menu.Component,
		Title:      menu.Title,
		Icon:       menu.Icon,
		Sort:       menu.Sort,
		Hiddent:    menu.Hidden,
		CreateTime: menu.CreateTime.Unix(),
		UpdateTime: menu.UpdateTime.Unix(),
	}
	if menu.DeleteTime.Valid {
		pbMenu.DeleteTime = menu.DeleteTime.Time.Unix()
	} else {
		pbMenu.DeleteTime = 0
	}
	return pbMenu, nil
}
