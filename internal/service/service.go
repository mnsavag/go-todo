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
}

type TodoItem interface {
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
