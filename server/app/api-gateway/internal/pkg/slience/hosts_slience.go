package slience

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"regexp"
	"time"
)

func AlarmIsMatchDefault(alarm types.Alerts, match map[string][]Sliences) (host string) {
	for host, consumerSliences := range match {
		for _, consumerMatcher := range consumerSliences {
			if consumerMatcher.IsDefault {
				if AlermIsMatch(consumerMatcher.Matchers, alarm) {
					return host
				}
			}
		}
	}

	return ""
}

func AlermIsMatch(matchers []types.Matchers, alarm types.Alerts) bool {
	for _, match := range matchers {
		if alarmvalue, ok := alarm.Labels[match.Name]; ok {
			if alarmValueIsMatch(match, alarmvalue) == false {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func alarmValueIsMatch(consumerMatch types.Matchers, value string) bool {
	if consumerMatch.IsRegex == true {
		re, err := regexp.Compile(consumerMatch.Value)
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
			return value == consumerMatch.Value
		} else /* env != prod */ {
			return !(value == consumerMatch.Value)
		}
	}
}

func GetConsumerSliences(db *gorm.DB) SafeSliences {
	slienceResults := []types.SlienceJoinRest{}
	err := db.Model(&types.Host{}).
		Select("`hosts`.`id`, `hosts`.`host`, `slience_names`.`default`, `slience_names`.`to`, `slience_names`.`slience_name`, `slience_matchers`.`name`, `slience_matchers`.`value`, `slience_matchers`.`is_regex`, `slience_matchers`.`is_equal`").
		Joins("JOIN slience_names ON hosts.id = slience_names.host_id").
		Joins("JOIN slience_matchers ON slience_names.id = slience_matchers.slience_name_id and slience_matchers.host_id = hosts.id").
		Scan(&slienceResults).Error
	if err != nil {
		panic(err)
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
		slienceResults[i].Value = templateValue.String()
		templateValue.Reset()
	}
	i, err := json.Marshal(slienceResults)
	logx.Infof("---> %s, error: %v", string(i), err)
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

	consumerSliences := SafeSliences{Sliences: make(map[string][]Sliences)}
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
					SliencesID:  "",
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
				SliencesID:  "",
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
	marshal, err := json.Marshal(consumerSliences.Sliences)
	logx.Infof("数据库静默规则: %s, error: %v", string(marshal), err)
	return consumerSliences
}

func AlertmanagerSliences(alertmanagerURL string, matchs []types.Matchers, host string, slienceName string) (silenceID string, err error) {
	now := time.Now()
	b := AMSlience{
		Matchers:  matchs,
		StartsAt:  now,
		EndsAt:    now.Add(time.Minute * 30),
		CreatedBy: "chaos root",
		Comment:   fmt.Sprintf("%s/%s", host, slienceName),
		ID:        "",
	}
	marshal, err := json.Marshal(b)
	if err != nil {
		return "", errors.Wrap(err, "body Marshal Failed")
	}
	body := bytes.NewBuffer(marshal)
	request, err := http.NewRequest(http.MethodPost, alertmanagerURL, body)
	if err != nil {
		return "", errors.Wrap(err, "create request Failed")
	}
	client := http.Client{Timeout: time.Second}
	response, err := client.Do(request)
	if err != nil {
		return "", errors.Wrap(err, "client exec Failed")
	}
	if response.StatusCode != 200 {
		return "", errors.New("alertmanager response code is not 200")
	}
	defer response.Body.Close()
	responseBody := AMSlienceReponse{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		return "", errors.Wrap(err, "alertmanager response body Decode Failed")
	}
	return responseBody.SilenceID, nil
}
