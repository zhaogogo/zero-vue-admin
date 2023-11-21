package slience

import (
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
	"time"
)

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

type AlertmanagerGetSlienceResponse []struct {
	ID        string     `json:"id"`
	Status    Status     `json:"status"`
	UpdatedAt time.Time  `json:"updatedAt"`
	Comment   string     `json:"comment"`
	CreatedBy string     `json:"createdBy"`
	EndsAt    time.Time  `json:"endsAt"`
	Matchers  []Matchers `json:"matchers"`
	StartsAt  time.Time  `json:"startsAt"`
}
type Status struct {
	State string `json:"state"`
}
type Matchers struct {
	IsEqual bool   `json:"isEqual"`
	IsRegex bool   `json:"isRegex"`
	Name    string `json:"name"`
	Value   string `json:"value"`
}
