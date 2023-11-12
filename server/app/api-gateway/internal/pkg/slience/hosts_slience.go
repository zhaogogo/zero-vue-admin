package slience

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/config"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func AlarmIsMatchDefault(alarm types.Alerts, match map[string][]Sliences) (host string) {
	fmt.Println("1--->", alarm.Labels)
	for host, consumerSliences := range match {
		for _, consumerSlience := range consumerSliences {
			if consumerSlience.IsDefault {
				if AlermIsMatch(consumerSlience.Matchers, alarm) {
					logx.Infof("默认匹配规则:%s-%s => %v 匹配 %v", host, consumerSlience.SlienceName, consumerSlience.Matchers, alarm.Labels)
					return host
				} else {
					logx.Infof("默认匹配规则:%s-%s => %v 不匹配 %v", host, consumerSlience.SlienceName, consumerSlience.Matchers, alarm.Labels)
				}
			}
		}
	}

	return ""
}

func AlermIsMatch(matchers []types.Matchers, alarm types.Alerts) bool {
	var res []bool
	for _, match := range matchers {
		if alarmvalue, ok := alarm.Labels[match.Name]; ok {
			if alarmValueIsMatch(match, alarmvalue) {
				//fmt.Println(">>>", match.Name, alarm.Labels[match.Name], alarmvalue, true)
				res = append(res, true)
			} else {
				//fmt.Println(">>>", match.Name, alarm.Labels[match.Name], alarmvalue, false)
				res = append(res, false)
			}
		}
	}
	for _, v := range res {
		if v == false {
			return false
		}
	}
	return true
}

func alarmValueIsMatch(consumerMatch types.Matchers, value string) bool {
	if consumerMatch.IsRegex == true {
		re, err := regexp.Compile(strings.Trim(consumerMatch.Value, "\""))
		if err != nil {
			logx.Errorf("正则compile错误: error: %v, value: %v", err, consumerMatch.Value)
			return false
		}
		if consumerMatch.IsEqual == true {
			return re.MatchString(value)
		} else {
			return !re.MatchString(value)
		}
	} else /* consumerMatch.IsRegex == false */ {
		if consumerMatch.IsEqual == true /* env = prod */ {
			//fmt.Println(value, consumerMatch.Value, value == consumerMatch.Value)
			return value == consumerMatch.Value
		} else /* env != prod */ {
			return !(value == consumerMatch.Value)
		}
	}
}

func GetConsumerSliences(db *gorm.DB, consumerSliences *SafeSliences) error {
	slienceResults := []types.SlienceJoinRest{}
	err := db.Model(&types.Host{}).
		Select("`hosts`.`id`, `hosts`.`host`, `slience_names`.`default`, `slience_names`.`to`, `slience_names`.`slience_name`, `slience_matchers`.`name`, `slience_matchers`.`value`, `slience_matchers`.`is_regex`, `slience_matchers`.`is_equal`").
		Joins("JOIN slience_names ON hosts.id = slience_names.host_id").
		Joins("JOIN slience_matchers ON slience_names.id = slience_matchers.slience_name_id and slience_matchers.host_id = hosts.id").
		Scan(&slienceResults).Error
	if err != nil {
		return errors.Wrapf(err, "slienceJoinRest 获取失败")
	}
	templateValue := bytes.NewBuffer(nil)
	for i, slienceResult := range slienceResults {
		parse, err := template.New("t1").Parse(slienceResult.Value)
		if err != nil {
			logx.Error(err)
			continue
		}
		err = parse.ExecuteTemplate(templateValue, "t1", slienceResult)
		if err != nil {
			logx.Error(err)
			continue
		}
		slienceResults[i].Value = strings.Trim(templateValue.String(), "\"")
		templateValue.Reset()
	}

	//sliences := SafeSliences{Sliences: make(map[string]map[string][]types.Matchers)}
	//sliences.Mu.Lock()
	//defer sliences.Mu.Unlock()
	//for _, res := range slienceResults {
	//	if sliences.Sliences[res.Host] == nil {
	//		sliences.Sliences[res.Host] = make(map[string][]types.Matchers)
	//	}
	//	if res.Default {
	//		sliences.Sliences[res.Host]["default"] = append(sliences.Sliences[res.Host]["default"], types.Matchers{
	//			Name:    res.Name,
	//			Value:   res.Value,
	//			IsRegex: false,
	//			IsEqual: true,
	//		})
	//	}
	//	sliences.Sliences[res.Host][res.SlienceName] = append(sliences.Sliences[res.Host][res.SlienceName], types.Matchers{
	//		Name:    res.Name,
	//		Value:   res.Value,
	//		IsRegex: false,
	//		IsEqual: true,
	//	})
	//}

	consumerSliences.Sliences = make(map[string][]Sliences)
	consumerSliences.Mu.Lock()
	defer consumerSliences.Mu.Unlock()

	for _, res := range slienceResults {
		if _, ok := consumerSliences.Sliences[res.Host]; ok {
			hasSlienceName := false
			for slienceIndex, slien := range consumerSliences.Sliences[res.Host] {
				if slien.SlienceName == res.SlienceName {
					consumerSliences.Sliences[res.Host][slienceIndex].Matchers = append(consumerSliences.Sliences[res.Host][slienceIndex].Matchers, types.Matchers{
						Name:    res.Name,
						Value:   res.Value,
						IsRegex: res.IsRegex,
						IsEqual: res.IsEqual,
					})
					hasSlienceName = true
				}
			}

			if !hasSlienceName {
				consumerSliences.Sliences[res.Host] = append(consumerSliences.Sliences[res.Host], Sliences{
					SlienceName: res.SlienceName,
					IsDefault:   res.Default,
					To:          res.To,
					Matchers: []types.Matchers{
						types.Matchers{
							Name:    res.Name,
							Value:   res.Value,
							IsRegex: res.IsRegex,
							IsEqual: res.IsEqual,
						},
					},
				})
			}

		} else {
			consumerSliences.Sliences[res.Host] = append(consumerSliences.Sliences[res.Host], Sliences{
				SlienceName: res.SlienceName,
				IsDefault:   res.Default,
				To:          res.To,
				Matchers: []types.Matchers{
					types.Matchers{
						Name:    res.Name,
						Value:   res.Value,
						IsRegex: res.IsRegex,
						IsEqual: res.IsEqual,
					},
				},
			})
		}
	}
	return nil
}

func AlertmanagerSliences(cfg config.MonitoringConfig, host string, duration string, slience Sliences) (silenceID string, err error) {
	now := time.Now()
	var alertmanagerUrl = ""
	if slience.To == 1 {
		alertmanagerUrl = cfg.AlertmanagerSlienceURL.ZW
	} else {
		alertmanagerUrl = cfg.AlertmanagerSlienceURL.YZ
	}
	var d time.Duration
	if duration == "" {
		duration = "30m"
	}
	d, err = time.ParseDuration(duration)
	if err != nil {
		return "", errors.Wrapf(err, "静默时间格式错误got: %s", duration)
	}
	b := AMSlience{
		Matchers:  slience.Matchers,
		StartsAt:  now,
		EndsAt:    now.Add(d),
		CreatedBy: "chaos root",
		Comment:   fmt.Sprintf("%s/%s", host, slience.SlienceName),
	}
	marshal, err := json.Marshal(b)
	if err != nil {
		return "", errors.Wrap(err, "body Marshal Failed")
	}
	body := bytes.NewBuffer(marshal)
	request, err := http.NewRequest(http.MethodPost, alertmanagerUrl+"/api/v2/silences", body)
	if err != nil {
		return "", errors.Wrap(err, "create request Failed")
	}
	request.Header.Add("Content-Type", "application/json")
	client := http.Client{Timeout: time.Second}
	response, err := client.Do(request)
	if err != nil {
		return "", errors.Wrap(err, "client exec Failed")
	}
	if response.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("alertmanager response code: got %d want 200", response.StatusCode))
	}
	defer response.Body.Close()
	responseBody := AMSlienceReponse{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		return "", errors.Wrap(err, "alertmanager response body Decode Failed")
	}
	return responseBody.SilenceID, nil
}

func AlertmanagerSliencesExpired(cfg config.MonitoringConfig, host string, slience Sliences) error {
	var alertmanagerUrl = ""
	if slience.To == 1 {
		alertmanagerUrl = cfg.AlertmanagerSlienceURL.ZW
	} else {
		alertmanagerUrl = cfg.AlertmanagerSlienceURL.YZ
	}
	sliencesResponse, err := http.Get(alertmanagerUrl + "/api/v2/silences")
	if err != nil {
		return errors.Wrapf(err, "获取alertmanager %s静默列表失败", alertmanagerUrl)
	}
	defer sliencesResponse.Body.Close()

	alertmanagerSliences := AlertmanagerGetSlienceResponse{}
	err = json.NewDecoder(sliencesResponse.Body).Decode(&alertmanagerSliences)
	if err != nil {
		return errors.Wrapf(err, "获取alertmanager %s静默响应json解析失败", alertmanagerUrl)
	}
	slienceID := []string{}
	for _, r := range alertmanagerSliences {
		if r.Comment == fmt.Sprintf("%s/%s", host, slience.SlienceName) && r.Status.State == "active" {
			slienceID = append(slienceID, r.ID)
		}
	}

	for _, id := range slienceID {
		request, err := http.NewRequest(http.MethodDelete, alertmanagerUrl+"/api/v2/silence/"+id, nil)
		if err != nil {
			return errors.Wrapf(err, "alertmanager静默过期NewRequest失败，url: %s/%s", alertmanagerUrl, id)
		}
		client := http.Client{}
		response, err := client.Do(request)
		if err != nil {
			return errors.Wrapf(err, "alertmanager静默过期Client.Do失败, url: %s/%s", alertmanagerUrl, id)
		}
		response.Body.Close()
	}
	return nil
}
