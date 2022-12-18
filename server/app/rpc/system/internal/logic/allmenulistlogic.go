package logic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AllMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAllMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllMenuListLogic {
	return &AllMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AllMenuListLogic) AllMenuList(in *pb.Empty) (*pb.AllMenuListResponse, error) {
	menus, err := l.svcCtx.MenuModel.FindAll_NC(l.ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "数据库查询失败")
	}
	pmenuList := []*pb.Menu{}
	for _, menu := range menus {
		pmenu := &pb.Menu{
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
			pmenu.DeleteTime = menu.DeleteTime.Time.Unix()
		} else {
			pmenu.DeleteTime = 0
		}
		pmenuList = append(pmenuList, pmenu)
	}
	return &pb.AllMenuListResponse{
		List: pmenuList,
	}, nil
}
