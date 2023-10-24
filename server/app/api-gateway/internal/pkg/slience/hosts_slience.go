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

func AlarmIsMatchDefault(alarm types.Alerts, match map[string]map[string][]types.Matchers) string {
	for host, consumerSliences := range match {
		consumerMatchs := consumerSliences["default"]
		isMatch := true

		for _ /* slience name */, consumerMatch := range consumerMatchs {
			if alarmvalue, ok := alarm.Labels[consumerMatch.Name]; ok {
				if alarmValueIsMatch(consumerMatch, alarmvalue) == false {
					isMatch = false
				}
			}
			if isMatch == false {
				return ""
			}
		}

		return host
	}

	return ""
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
	db.Model(&types.Host{}).
		Select("`hosts`.`id`,`hosts`.`host`, `slience_names`.`default`,`slience_names`.`slience_name`, `slience_matchers`.`name`, `slience_matchers`.`value`, `slience_matchers`.`is_regex`, `slience_matchers`.`is_equal`").
		Joins("JOIN slience_names ON hosts.id = slience_names.host_id").
		Joins("JOIN slience_matchers ON slience_names.id = slience_matchers.slience_name_id and and slience_matchers.host_id = hosts.id").
		Scan(&slienceResults)

	templateValue := bytes.NewBuffer(nil)
	for i, slienceResult := range slienceResults {
		parse, err := template.New("t1").Parse(slienceResult.Value)
		if err != nil {
			logx.Error(err)
			continue
		}
		err = parse.ExecuteTemplate(templateValue, "t1", parse)
		if err != nil {
			logx.Error(err)
			continue
		}
		slienceResults[i].Value = templateValue.String()
		templateValue.Reset()
	}
	sliences := SafeSliences{Sliences: make(map[string]map[string][]types.Matchers)}
	sliences.Mu.Lock()
	defer sliences.Mu.Unlock()
	for _, res := range slienceResults {
		if sliences.Sliences[res.Host] == nil {
			sliences.Sliences[res.Host] = make(map[string][]types.Matchers)
		}
		if res.Default {
			sliences.Sliences[res.Host]["default"] = append(sliences.Sliences[res.Host]["default"], types.Matchers{
				Name:    res.Name,
				Value:   res.Value,
				IsRegex: false,
				IsEqual: true,
			})
		}
		sliences.Sliences[res.Host][res.SlienceName] = append(sliences.Sliences[res.Host][res.SlienceName], types.Matchers{
			Name:    res.Name,
			Value:   res.Value,
			IsRegex: false,
			IsEqual: true,
		})
	}
	return sliences
}

func AlertmanagerSliences(alertmanagerURL string, matchs []types.Matchers, slienceName string) error {
	now := time.Now()
	b := AMSlience{
		Matchers:  matchs,
		StartsAt:  now,
		EndsAt:    now.Add(time.Hour * 2),
		CreatedBy: "chaos root",
		Comment:   fmt.Sprintf("chaos slience => [%s]", slienceName),
		ID:        "",
	}
	marshal, err := json.Marshal(b)
	if err != nil {
		return errors.Wrap(err, "body Marshal Failed")
	}
	body := bytes.NewBuffer(marshal)
	request, err := http.NewRequest(http.MethodPost, alertmanagerURL, body)
	if err != nil {
		return errors.Wrap(err, "create request Failed")
	}
	client := http.Client{Timeout: time.Second}
	response, err := client.Do(request)
	if err != nil {
		return errors.Wrap(err, "client exec Failed")
	}
	if response.StatusCode != 200 {
		return errors.New("alertmanager response code is not 200")
	}
	defer response.Body.Close()
	responseBody := AMSlienceReponse{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		return errors.Wrap(err, "alertmanager response body Decode Failed")
	}
	return nil
}
