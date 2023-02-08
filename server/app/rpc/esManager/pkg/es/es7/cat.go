package es7

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/pkg/es"
)

type cat struct {
	client *elastic.Client
	err    error
}

func (c *cat) Health(ctx context.Context, param es.CatParam) (catrest interface{}, err error) {
	if c.err != nil {
		return nil, errors.Wrap(err, "创建客户端失败")
	}

	resp, err := c.client.CatHealth().
		Pretty(param.Pretty).
		Human(param.Human).
		ErrorTrace(param.ErrorTrace).
		FilterPath(param.FilterPath...).
		Do(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "执行错误")
	}
	r, err := json.Marshal(resp)
	if err != nil {
		return nil, errors.Wrap(err, "编码json字符串错误")
	}
	return r, nil
}
