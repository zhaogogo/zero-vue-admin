package slience

import (
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
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
