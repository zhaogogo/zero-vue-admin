package hosts

import (
	"context"
	"fmt"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/slience"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandlerHostsSlienceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandlerHostsSlienceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandlerHostsSlienceLogic {
	return &HandlerHostsSlienceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandlerHostsSlienceLogic) HandlerHostsSlience(req *types.HandlerHostsSlienceRequest) (resp *types.HttpCommonResponse, err error) {
	l.svcCtx.SlienceList.Mu.RLock()
	slienceList := l.svcCtx.SlienceList.Sliences
	l.svcCtx.SlienceList.Mu.RUnlock()
	errRes := map[string][]string{}
	for _, host := range req.Hosts {
		if slienceNames, ok := slienceList[host]; ok {
			if req.OpType == "active" {
				slienceNamesed := []string{}
				for _, s := range slienceNames {
					_, err := slience.AlertmanagerSliences(l.svcCtx.Config.MonitoringConfig, host, req.Duration, s)
					if err != nil {
						logx.Errorf("调用alertmanager API静默失败, host: %s, slience_name: %s, error: %v", host, s.SlienceName, err)
						errRes[host] = append(errRes[host], s.SlienceName)
					} else {
						slienceNamesed = append(slienceNamesed, s.SlienceName)
					}
				}
				logx.Infof("主机%s静默%v", host, slienceNamesed)
			} else {
				for _, s := range slienceNames {
					err := slience.AlertmanagerSliencesExpired(l.svcCtx.Config.MonitoringConfig, host, s)
					if err != nil {
						return nil, errorx.New(err, "静默过期失败")
					}
				}
			}
		} else {
			logx.Errorf("hosts (%s) slience failed", host)
			errRes[host] = append(errRes[host], fmt.Sprintf("slienceList 无此主机%s", host))
		}
	}
	var res *types.HttpCommonResponse
	if len(errRes) == 0 {
		res = &types.HttpCommonResponse{
			Code: 200,
			Msg:  "OK",
			Meta: nil,
		}
	} else {
		res = &types.HttpCommonResponse{
			Code: 200,
			Msg:  fmt.Sprintf("%v", errRes),
			Meta: nil,
		}
	}

	return res, nil
}
