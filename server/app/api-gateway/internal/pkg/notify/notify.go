package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
	"net/http"
	"time"
)

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

func SentMessage(u string, slience types.SlienceName) {

}
