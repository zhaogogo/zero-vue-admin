package es7

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/pkg/es"
	"net/http"
	"strings"

	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
)

type es7 struct {
	CacheID  uint64
	URL      string
	User     string
	PassWord string
}

func NewES7(url string, user string, password string, id uint64) *es7 {
	return &es7{
		CacheID:  id,
		URL:      url,
		User:     user,
		PassWord: password,
	}
}

func (e *es7) getClient() (*elastic.Client, error) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	optList := []elastic.ClientOptionFunc{
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		elastic.SetURL(strings.Split(e.URL, ",")...),
		elastic.SetHttpClient(httpClient),
		elastic.SetBasicAuth(e.User, e.PassWord),
	}
	client, err := elastic.NewSimpleClient(optList...)
	if err != nil {
		return nil, errors.Wrap(err, "ES创建客户端失败")
	}
	return client, nil
}

func (e *es7) Ping(ctx context.Context) (pingres interface{}, err error) {
	client, err := e.getClient()
	if err != nil {
		return nil, err
	}
	pingRes := &elastic.PingResult{}
	for _, u := range strings.Split(e.URL, ",") {
		pingRes, _, err = client.Ping(u).Do(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "client ping 失败")
		}
		if pingRes.Version.Number == "" {
			return nil, errors.New(fmt.Sprintf("ES URL: %s OK,账号密码错误", u))
		}
	}

	return pingRes, nil
}

func (e *es7) Cat() es.CatInterface {
	client, err := e.getClient()

	return &cat{
		client: client,
		err:    err,
	}
}
