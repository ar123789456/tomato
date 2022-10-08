package models

type User struct {
	Name       string  `json:"name"`
	SecondName *string `json:"secondName"`
	Nick       string  `json:"nick"`
	Email      *string `json:"email"`
	Photo      *string `json:"photo"`
	Class      string  `json:"class"`
}
