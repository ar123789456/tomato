package models

import (
	"github.com/google/uuid"
)

type Tomato struct {
	Id             uuid.UUID      `json:"id"`
	TimerStart     int64          `json:"timerStart"`
	CreateTime     int64          `json:"createTime"`
	CategoryTomato CategoryTomato `json:"categoryTomato"`
	Title          string         `json:"title"`
	Context        string         `json:"context"`
	Tags           []string       `json:"tags"`
}

type CategoryTomato struct {
	WorkTime int64  `json:"workTime"` //minute
	Rest     *int64 `json:"rest"`     //minute
}

func (ct *CategoryTomato) SetRest() {
	if ct.Rest == nil {
		rest := ct.WorkTime / 6
		ct.Rest = &rest
	}
}
