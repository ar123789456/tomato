package models

import "github.com/google/uuid"

//Habit model
type Habit struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsDeleted   bool      `json:"-"`
	Completed   []bool    `json:"completed"`
}
