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

func (s *TodoListService) GetAll(userId int64) ([]model.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int64) (model.TodoList, error) {
	return s.repo.GetById(userId, listId)
}
