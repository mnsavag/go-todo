package service

import (
	"goTodo/internal/model"
	"goTodo/internal/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int64, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int64, error)
}

type TodoList interface {
	Create(userId int64, list model.TodoList) (int64, error)
	GetAll(userId int64) ([]model.TodoList, error)
	GetById(userId, listId int64) (model.TodoList, error)
	Delete(userId, listId int64) error
	Update(userId, listId int64, input model.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int64, item model.TodoItem) (int64, error)
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
