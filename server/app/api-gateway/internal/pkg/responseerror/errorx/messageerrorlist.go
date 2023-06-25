package errorx

import (
	"sync"
)

type MsgErrList struct {
	List []interface{}
	mu   sync.Mutex
}

func (e *MsgErrList) WithMeta(rpcMethod string, rpcerr string, param interface{}) *MsgErrList {
	e.mu.Lock()
	defer e.mu.Unlock()
	if rpcMethod == "" && rpcerr == "" {
		e.List = append(e.List, param)
		return e
	}
	e.List = append(
		e.List,
		map[string]interface{}{
			"rpc方法": rpcMethod,
			"rpc错误": rpcerr, //如果rocerr类型是error json marshal
			"请求参数":  param,
		},
	)
	return e
}
