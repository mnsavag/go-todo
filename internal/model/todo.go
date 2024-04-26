package model

type TodoList struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoItem struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description int    `json:"description"`
	Done        bool   `json:"done"`
}

type ListItem struct {
	Id     int64
	ListId int
	ItemId int
}

type UsersList struct {
	Id     int64
	UserId int
	ListId int
}
