package repository

import (
	"database/sql"

	"goTodo/internal/model"
	"goTodo/internal/repository/sqlite"
)

type Authorization interface {
	CreateUser(user model.User) (int64, error)
	GetRegisteredUser(username, password string) (model.User, error)
}

type TodoList interface {
	Create(userId int64, list model.TodoList) (int64, error)
	GetAll(userId int64) ([]model.TodoList, error)
	GetById(userId, listId int64) (model.TodoList, error)
	Delete(userId, listId int64) error
	Update(userId, listId int64, input model.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int64, item model.TodoItem) (int64, error)
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: sqlite.NewAuthSqlite(db),
		TodoList:      sqlite.NewTodoListSqlite(db),
	}
}
