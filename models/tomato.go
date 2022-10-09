package models

import (
	"github.com/google/uuid"
)

type Tomato struct {
	Id             uuid.UUID   `json:"id"`
	TimerStart     int64       `json:"timerStart"`
	CreateTime     int64       `json:"createTime"`
	CategoryTomato TimerTomato `json:"categoryTomato"`
	Title          string      `json:"title"`
	Context        *string     `json:"context"`
	Tags           []string    `json:"tags"`
}

type TimerTomato struct {
	WorkTime int64  `json:"workTime"` //minute
	Rest     *int64 `json:"rest"`     //minute
}

func (tt *TimerTomato) SetRest() {
	if tt.Rest == nil {
		rest := tt.WorkTime / 6
		tt.Rest = &rest
	}
}

//-------------------------------------
// Analytics
//-------------------------------------

type TomatoNltx struct {
}

//-------------------------------------
// Input
//-------------------------------------

//CreateTomatoIn input for create
type CreateTomatoIn struct {
	TimerStart     int64       `json:"timerStart"`
	CategoryTomato TimerTomato `json:"categoryTomato"`
	Title          string      `json:"title"`
	Context        *string     `json:"context"`
	Tags           []string    `json:"tags"`
}

//TomatoIn input value Start and delete
type TomatoIn struct {
	Id uuid.UUID `json:"id"`
}
