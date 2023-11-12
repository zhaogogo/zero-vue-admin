package alarm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/chanhandle"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/slience"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
	"net/http"
	"strings"
	"time"

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
	l.svcCtx.SlienceList.Mu.RLock()
	consumerMatch := l.svcCtx.SlienceList.Sliences
	l.svcCtx.SlienceList.Mu.RUnlock()
	marshal, err := json.MarshalIndent(req, "", "\t")
	fmt.Println(string(marshal), err)
	for _, alert := range req.Alerts {
		if host := slience.AlarmIsMatchDefault(alert, consumerMatch); host != "" {
			if alert.Status == "firing" {
				slienceNames := []string{}
				for _, sli := range consumerMatch[host] {
					slienceNames = append(slienceNames, sli.SlienceName)
					_, err := slience.AlertmanagerSliences(l.svcCtx.Config.MonitoringConfig, host, "", sli)
					if err != nil {
						logx.Errorf("调用alertmanager API静默失败, host: %s, slience_name: %s, error: %v", host, sli.SlienceName, err)
					}
				}
				var a = Alert{URL: "http://127.0.0.1:8075/api/v2/idatas"}
				message, err := a.SentMessage(&AlertMessage{
					Title:      fmt.Sprintf("k8s关联静默"),
					Message:    strings.Join(slienceNames, "\n,"),
					NoticeName: fmt.Sprintf("%s", host),
					Serverity:  "P2",
				})
				if err != nil {
					logx.Error("发送消息失败, host: %s, body: %s", host, message)
				}
			} else {

			}
		}

	}
	return nil
}

// TiDB生产集群出现执行时间大于20秒的SQL：
// tidb实例：instance
// 客户端：host
// 库名：db
// 账号：user
// 执行时间：time
// SQL内容：sql查询结果的sqltxt

//	{
//		"notice_name": "test2",
//		"serverity": "P0",
//		"summary": "test_summary",
//		"message": "test_message"
//	}
type AlertMessage struct {
	Message    string `json:"message"`
	NoticeName string `json:"notice_name"`
	Serverity  string `json:"serverity"`
	Title      string `json:"summary"`
	DBName     string `json:"-"`
}

// test_summary
// test2
// 2022-09-28 18:13:08<br/>通知消息：test_message

type Alert struct {
	URL string `yaml:"url"`
}

func (u *Alert) SentMessage(data *AlertMessage) (string, error) {
	j, err := json.Marshal(data)
	if err != nil {
		return "", errors.Wrap(err, "json encode error")
	}
	body := bytes.NewBuffer(j)

	alertURL := u.URL

	req, err := http.NewRequest("POST", alertURL, body)
	if err != nil {
		return string(j), errors.Wrap(err, "new request error")
	}
	client := &http.Client{
		Timeout: time.Second * 2,
	}
	response, err := client.Do(req)
	if err != nil {
		return string(j), errors.Wrap(err, "http client do error")
	}
	logx.Infof("Send message response HTTP StatusCode is: %v", response.StatusCode)
	var res = struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}{}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return string(j), errors.Wrap(err, "response body json decoder error")
	}
	if res.Code != 200 {
		return string(j), errors.New(fmt.Sprintf("response body code is got %d, msg: %s", res.Code, res.Msg))
	}
	return string(j), nil
}

type MatcherDefault struct {
	chanhandle.Next
}

func (m *MatcherDefault) Do(s *chanhandle.SlienceChan) error {

	return nil
}
