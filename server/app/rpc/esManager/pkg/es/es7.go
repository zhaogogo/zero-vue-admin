package es

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

type es7 struct {
	URL      string
	User     string
	PassWord string
}

func (e *es7) GetClient() (*elastic.Client, error) {
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
	return elastic.NewSimpleClient(optList...)
}

func (e *es7) Ping() (res interface{}, err error) {
	client, err := e.GetClient()
	if err != nil {
		return nil, err
	}
	pingRes := &elastic.PingResult{}
	for _, u := range strings.Split(e.URL, ",") {
		pingRes, _, err = client.Ping(u).Do(context.Background())
		if err != nil {
			return nil, err
		}
		if pingRes.Version.Number == "" {
			return nil, errors.New(fmt.Sprintf("ES URL: %s OK,密码错误", u))
		}
	}
	return pingRes, nil
}
