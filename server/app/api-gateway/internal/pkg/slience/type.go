package slience

import (
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
	"sync"
	"time"
)

type SafeSliences struct {
	Mu sync.RWMutex
	//           instance   alertName
	Sliences map[string]map[string][]types.Matchers
	//Sliences map[string]map[string]Slience
}

type Slience struct {
	To       int
	Matchers []types.Matchers
}

type AMSlience struct {
	Matchers  []types.Matchers `json:"matchers"`
	StartsAt  time.Time        `json:"startsAt"`
	EndsAt    time.Time        `json:"endsAt"`
	CreatedBy string           `json:"createdBy"`
	Comment   string           `json:"comment"`
	ID        string           `json:"id"`
}

type AMSlienceReponse struct {
	SilenceID string `json:"silenceID"`
}
