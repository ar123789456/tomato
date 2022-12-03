package models

import (
	"github.com/google/uuid"
)

// User model
type User struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	SecondName *string   `json:"secondName"`
	Nick       string    `json:"nick"`
	Email      *string   `json:"email"`
	Password   string    `json:"password"`
}
