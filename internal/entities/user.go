package entities

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Username int    `json:"username"`
	Password int    `json:"password"`
}
