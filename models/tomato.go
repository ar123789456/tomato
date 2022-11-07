package models

import (
	"errors"
	"github.com/google/uuid"
	"strings"
)

type Tomato struct {
	Id          uuid.UUID   `json:"id"`
	TimerStart  int64       `json:"timerStart"`
	CreateTime  int64       `json:"createTime"`
	TimerTomato TimerTomato `json:"categoryTomato"`
	Title       string      `json:"title"`
	Context     *string     `json:"context"`
	Tags        []string    `json:"tags"`
	UserID      uuid.UUID   `json:"-"`
}

var (
	ErrInvalidTitle = errors.New("invalid title")
)

func (t *Tomato) Validate() error {
	if t.Id == uuid.Nil {
		return ErrInvalidId
	}
	if strings.TrimSpace(t.Title) == "" {
		return ErrInvalidTitle
	}
	return t.TimerTomato.Validate()
}

type TimerTomato struct {
	id       int64  `json:"-"`
	WorkTime int64  `json:"workTime"` //minute
	Rest     *int64 `json:"rest"`     //minute
}

var ErrInvalidWorkTime = errors.New("invalid work time")

func (tt *TimerTomato) Validate() error {
	if tt.WorkTime == 0 {
		return ErrInvalidWorkTime
	}
	return nil
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
	TimerStart  int64       `json:"timerStart"`
	TimerTomato TimerTomato `json:"categoryTomato"`
	Title       string      `json:"title"`
	Context     *string     `json:"context"`
	Tags        []string    `json:"tags"`
}

//TomatoIn input value Start and delete
type TomatoIn struct {
	Id uuid.UUID `json:"id"`
}
