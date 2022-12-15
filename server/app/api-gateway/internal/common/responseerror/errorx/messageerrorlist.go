package errorx

import (
	"sync"
)

type MsgErrList struct {
	List []string
	mu   sync.Mutex
}

func (m *MsgErrList) Append(v string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.List = append(m.List, v)
}
