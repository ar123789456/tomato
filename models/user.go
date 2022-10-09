package models

import "github.com/google/uuid"

type User struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	SecondName *string   `json:"secondName"`
	Nick       string    `json:"nick"`
	Email      *string   `json:"email"`
	Photo      *string   `json:"photo"`
	Class      string    `json:"class"`
}
