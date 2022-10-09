package models

import (
	"errors"
	"github.com/google/uuid"
	"net/mail"
	"strings"
)

type User struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	SecondName *string   `json:"secondName"`
	Nick       string    `json:"nick"`
	Email      *string   `json:"email"`
	Photo      *string   `json:"photo"`
	Class      string    `json:"class"`
}

var (
	ErrInvalidName  = errors.New("invalid name")
	ErrInvalidNick  = errors.New("invalid nickname")
	ErrInvalidId    = errors.New("invalid id")
	ErrInvalidEmail = errors.New("invalid email")
	ErrInvalidClass = errors.New("invalid class")
)

func (u *User) Validate() error {
	if strings.TrimSpace(u.Name) == "" {
		return ErrInvalidName
	}
	if len(u.Name) == 0 || len(u.Name) < 5 || len(u.Name) > 15 {
		return ErrInvalidName
	}
	if len(u.Nick) == 0 || len(u.Nick) < 5 || len(u.Nick) > 15 {
		return ErrInvalidNick
	}
	if u.Id == uuid.Nil {
		return ErrInvalidId
	}
	if u.Email != nil {
		if _, err := mail.ParseAddress(*u.Email); err != nil {
			return ErrInvalidEmail
		}
	}
	if strings.TrimSpace(u.Class) == "" {
		return ErrInvalidClass
	}

	return nil
}

//-------------------------------------
// Input
//-------------------------------------

type CreateUser struct {
	Name       string  `json:"name"`
	SecondName *string `json:"secondName"`
	Nick       string  `json:"nick"`
	Email      *string `json:"email"`
	Photo      *string `json:"photo"`
	Class      string  `json:"class"`
}

type EditUser struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	SecondName *string   `json:"secondName"`
	Nick       string    `json:"nick"`
	Email      *string   `json:"email"`
	Photo      *string   `json:"photo"`
	Class      string    `json:"class"`
}
