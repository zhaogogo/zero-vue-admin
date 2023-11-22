package slience

import (
	"bytes"
	"fmt"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
	"html/template"
	"testing"
)

type Entry struct {
	Data types.Alerts
	Want bool
}

func TestAlermIsMatch(t *testing.T) {
	s := Sliences{
		SlienceName: "主机ping",
		IsDefault:   true,
		To:          0,
		Matchers: []types.Matchers{
			{
				Name:    "alertname",
				Value:   "Probe_PING_Node_Down_Prod",
				IsRegex: false,
				IsEqual: true,
			},
			{
				Name:    "instance",
				Value:   "10.100.124.101",
				IsRegex: false,
				IsEqual: true,
			},
		},
	}
	alerts := []Entry{
		{
			Data: types.Alerts{Labels: map[string]string{"alertname": "Probe_PING_Node_Down_Prod", "instance": "10.100.124.101"}},
			Want: true,
		},
		{
			Data: types.Alerts{Labels: map[string]string{"alertname": "Probe_PING_Node_Down_Prod111", "instance": "10.100.124.101"}},
			Want: false,
		},
		{
			Data: types.Alerts{Labels: map[string]string{"alertname": "Probe_PING_Node_Down_Prod", "instance": "10.100.124.101123"}},
			Want: false,
		},
		{
			Data: types.Alerts{Labels: map[string]string{"alertname": "Probe_PING_Node_Down_Prod", "instance": "10.100.124.101", "env": "prod"}},
			Want: true,
		},
		{
			Data: types.Alerts{Labels: map[string]string{"alertname": "Probe_PING_Node_Down_Prod22", "instance": "10.100.124.101", "env": "prod"}},
			Want: false,
		},
		{
			Data: types.Alerts{Labels: map[string]string{"alertname": "Probe_PING_Node_Down_Prod"}},
			Want: false,
		},
	}
	for _, alert := range alerts {
		match := AlermIsMatch(s, alert.Data)
		t.Logf("%t %#v\n", match, alert.Data.Labels)
		if match != alert.Want {
			t.Errorf("entry: %#v, 不匹配", alert)
		}
	}
}

type TT struct {
	SilenceNameDefault     string
	Silences               []Sliences
	SendAlertmanagerFailed []string
}

func TestNotifyTemplate(t *testing.T) {
	path := `C:\Users\qiang.zhao\Desktop\owner\zero-vue-admin\server\app\api-gateway\etc\t.template`
	templateData := TT{
		SilenceNameDefault: "Probe_PING_Node_Down_Prod",
		Silences: []Sliences{
			{
				SlienceName: "Probe_PING_Node_Down_Prod",
				IsDefault:   true,
				To:          1,
				Matchers: []types.Matchers{
					{
						Name:    "alertname",
						Value:   "Probe_PING_Node_Down_Prod",
						IsRegex: false,
						IsEqual: true,
					},
					{
						Name:    "instance",
						Value:   "192.168.14.101",
						IsRegex: false,
						IsEqual: true,
					},
				},
			},
			{
				SlienceName: "NodeDownProd",
				IsDefault:   false,
				To:          1,
				Matchers: []types.Matchers{
					{
						Name:    "alertname",
						Value:   "NodeDownProd",
						IsRegex: false,
						IsEqual: true,
					},
					{
						Name:    "instance",
						Value:   "192.168.14.101",
						IsRegex: false,
						IsEqual: true,
					},
				},
			},
			{
				SlienceName: "DockerdProcessDown",
				IsDefault:   false,
				To:          1,
				Matchers: []types.Matchers{
					{
						Name:    "alertname",
						Value:   "DockerdProcessDown",
						IsRegex: false,
						IsEqual: true,
					},
					{
						Name:    "instance",
						Value:   "192.168.14.101",
						IsRegex: false,
						IsEqual: true,
					},
				},
			},
			{
				SlienceName: "NodeStatus",
				IsDefault:   false,
				To:          1,
				Matchers: []types.Matchers{
					{
						Name:    "alertname",
						Value:   "NodeStatus",
						IsRegex: false,
						IsEqual: true,
					},
					{
						Name:    "instance",
						Value:   "192.168.14.101",
						IsRegex: false,
						IsEqual: true,
					},
				},
			},
		},
		SendAlertmanagerFailed: []string{"NodeDownProd"},
	}
	titleBuf := bytes.NewBuffer(nil)
	msgBuf := bytes.NewBuffer(nil)
	tpl, err := template.New("notify").ParseFiles(path)
	if err != nil {
		t.Error(err)
		return
	}
	err = tpl.ExecuteTemplate(titleBuf, "title", "")
	if err != nil {
		t.Error(err)
		return
	}
	err = tpl.ExecuteTemplate(msgBuf, "message", templateData)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("====================")
	fmt.Println("title:", titleBuf.String())
	fmt.Println("====================")
	fmt.Println(msgBuf.String())
}
