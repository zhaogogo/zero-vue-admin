package alarm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
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
	//marshal, err := json.MarshalIndent(req, "", "\t")
	//fmt.Println(string(marshal), err)
	for _, alert := range req.Alerts {
		if host, silenceNameDefault := slience.AlarmIsMatchDefault(alert, consumerMatch); host != "" {
			if alert.Status == "firing" {
				alertManagerCallFailedErrors := []error{}
				alertManagerCallFailedSilenceNames := []string{}
				for _, silence := range consumerMatch[host] {
					_, err := slience.AlertmanagerSliences(l.svcCtx.Config.MonitoringConfig, host, "", silence)
					if err != nil {
						alertManagerCallFailedErrors = append(alertManagerCallFailedErrors, &CallAlertmanagerError{
							Host:    host,
							Slience: silence,
							Root:    err,
						})
						alertManagerCallFailedSilenceNames = append(alertManagerCallFailedSilenceNames, silence.SlienceName)
					}
					zlog.Info().Str("active", "callAlertamanger").Dict("silence", zerolog.Dict().Str("host", host).Str("silencename", silence.SlienceName).Str("matchers", silence.StringMatchers())).Send()
				}
				if len(alertManagerCallFailedErrors) != 0 {
					zlog.Error().Errs("errors", alertManagerCallFailedErrors).Send()
				}
				var a = Alert{URL: l.svcCtx.Config.MonitoringConfig.NotifyURL}
				var (
					title string
					msg   string
				)
				titleBuf := bytes.NewBuffer(nil)
				msgBuf := bytes.NewBuffer(nil)
				t := l.svcCtx.NotifyTemplate
				if err := t.ExecuteTemplate(titleBuf, "title", ""); err == nil {
					title = strings.Trim(strings.TrimSpace(fmt.Sprintf("%s %s", titleBuf.String(), host)), "\"")
				} else {
					title = fmt.Sprintf("关联告警 %s (Default)", host)
					zlog.Error().Msgf("模板生成title失败 host: %s, error: %v", host, err)
				}
				if err := t.ExecuteTemplate(msgBuf, "message", struct {
					SilenceNameDefault     string
					Host                   string
					Silences               []slience.Sliences
					SendAlertmanagerFailed []string
				}{
					silenceNameDefault,
					host,
					consumerMatch[host],
					alertManagerCallFailedSilenceNames},
				); err == nil {
					msg = fmt.Sprintf("\n%s\n", strings.Trim(strings.TrimSpace(msgBuf.String()), "\""))
				} else {
					s := []string{}
					for _, silenceNames := range consumerMatch[host] {
						for _, failedSilenceName := range alertManagerCallFailedSilenceNames {
							if failedSilenceName == silenceNames.SlienceName {
								continue
							}
						}
						s = append(s, silenceNames.SlienceName)
					}
					msg = fmt.Sprintf("触发规则: %s\n关联静默: \n%s", silenceNameDefault, strings.Join(s, ",\n"))
					zlog.Error().Msgf("模板生成message失败 host: %s, error: %v", host, err)
				}

				request, err := a.SentMessage(&AlertMessage{
					Title:      title,
					Message:    msg,
					NoticeName: l.svcCtx.Config.MonitoringConfig.AggregationNotify,
					Serverity:  l.svcCtx.Config.MonitoringConfig.AggregationSeverity,
				})
				if err != nil {
					zlog.Error().Err(err).Str("body", request).Msgf("发送消息失败, host: %s", host)
				} else {
					zlog.Info().Str("body", request).Msg("发送消息成功")
				}
			} else {
				// 告警恢复逻辑
			}
		}

	}
	return nil
}

type CallAlertmanagerError struct {
	Host    string
	Slience slience.Sliences
	Root    error
}

func (e *CallAlertmanagerError) Error() string {
	return fmt.Sprintf("host: %s silence_name:%s matchers: %s root_error: %v", e.Host, e.Slience.SlienceName, e.Slience.StringMatchers(), e.Root)
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

func (u *Alert) SentMessage(data *AlertMessage) (request string, err error) {
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
	zlog.Info().Msgf("Send message response HTTP StatusCode is: %v", response.StatusCode)
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
