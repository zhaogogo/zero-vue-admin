package file

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/model/monitoring"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EnvSelectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEnvSelectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnvSelectLogic {
	return &EnvSelectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EnvSelectLogic) EnvSelect() (resp *types.EnvSelectResponseResponse, err error) {
	conns := []monitoring.StoreConnectManager{}
	if err = l.svcCtx.MonitoringDB().Find(&conns).Error; err != nil {
		return nil, errorx.New(err, "获取环境失败")
	}
	res := []types.Options{}
	envs := map[string]int{}
	envMap := map[string]string{"zw": "兆维", "yz": "亦庄"}
	for _, conn := range conns {
		envs[envMap[conn.Env]] += 1
	}
	for env := range envMap {
		r := types.Options{}
		if env == "兆维" {
			r.Label = "zw"
		} else {
			r.Label = "yz"
		}

		r.Value = env
		res = append(res, r)
	}
	for i, v := range res {
		c := []types.Options{}
		for _, conn := range conns {
			if v.Value == conn.Env {
				r := types.Options{}
				r.Value = conn.Host
				r.Label = conn.Host
				c = append(c, r)
			}
		}
		res[i].Children = c
	}

	return &types.EnvSelectResponseResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK", Meta: nil},
		List:               res,
	}, nil
}

func AAA(conns []monitoring.StoreConnectManager, res *types.Options) *types.Options {

	return res
}
