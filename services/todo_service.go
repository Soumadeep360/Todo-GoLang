package services

import (
	"errors"
	"todo-app/models"
	"todo-app/repository"
)

type TodoService struct {
	repo *repository.TodoRepository
}

func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) CreateTodo(todo *models.Todo) error {
	if todo.Title == "" {
		return errors.New("title is required")
	}
	return s.repo.Create(todo)
}

func (s *TodoService) GetTodos() ([]models.Todo, error) {
	return s.repo.GetAll()
}

func (s *TodoService) GetTodoByID(id uint) (*models.Todo, error) {
	return s.repo.GetByID(id)
}

func (s *TodoService) UpdateTodo(todo *models.Todo) error {
	return s.repo.Update(todo)
}

func (s *TodoService) DeleteTodo(id uint) error {
	return s.repo.Delete(id)
}
