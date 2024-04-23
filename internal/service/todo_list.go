package service

import (
	"goTodo/internal/model"
	"goTodo/internal/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int64, list model.TodoList) (int64, error) {
	return s.repo.Create(userId, list)
}
