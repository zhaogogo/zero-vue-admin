package alarm

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/slience"
	"strings"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WebhookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWebhookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WebhookLogic {
	return &WebhookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WebhookLogic) Webhook(req *types.AlarmRequest) error {
	machineRoom := map[string]string{
		"1": "http://192.168.14.105:9091/api/v1/sliences",
		"2": "http://10.100.114.105:9091/api/v1/sliences",
	}
	l.svcCtx.SlienceList.Mu.RLock()
	matchDefault := l.svcCtx.SlienceList.Sliences
	defer l.svcCtx.SlienceList.Mu.RUnlock()

	for _, alert := range req.Alerts {
		if host := slience.AlarmIsMatchDefault(alert, matchDefault); host != "" {
			if alert.Status == "firing" {
				for slienceName, matchs := range l.svcCtx.SlienceList.Sliences[host] {
					slienceto := strings.SplitN(slienceName, ":", 2)
					if len(slienceto) != 2 {
						logx.Errorf("获取机房位置失败, host: %s, slience_name: %s", host, slienceName)
						continue
					}
					if err := slience.AlertmanagerSliences(machineRoom[slienceto[1]], matchs, slienceName); err != nil {
						logx.Error(err)
						continue
					}
				}
			} else {

			}
		}

	}
	return nil
}
