package models

import "github.com/google/uuid"

// Task model
type Task struct {
	Id          uuid.UUID `json:"id"`
	CreatedAt   int64     `json:"createdAt"` //milliseconds
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserId      uuid.UUID `json:"-"`
	Completed   bool      `json:"completed"`
}
