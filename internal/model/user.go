package model

type User struct {
	Id       int64  `json:"id" db:"id"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
