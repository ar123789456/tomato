package models

import (
	"github.com/google/uuid"
)

// Tomato model
type Tomato struct {
	Id         uuid.UUID `json:"id"`
	TimerStart int64     `json:"timerStart"`
	CreateTime int64     `json:"createTime"`
	TimerId    uuid.UUID `json:"categoryTomato"` //Timer Id
	TaskId     uuid.UUID `json:"taskId"`         //Task Id
	UserId     uuid.UUID `json:"-"`              //User Id
}

// Timer model
type Timer struct {
	Id       int64     `json:"-"`
	UserId   uuid.UUID `json:"-"`        //User Id
	WorkTime int64     `json:"workTime"` //milliseconds
	Rest     *int64    `json:"rest"`     //milliseconds
}
