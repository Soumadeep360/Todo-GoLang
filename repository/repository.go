package repository

import (
	"todo-app/database"
	"todo-app/models"
)

type TodoRepository struct{}

func (r *TodoRepository) Create(todo *models.Todo) error {
	return database.DB.Create(todo).Error
}

func (r *TodoRepository) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	err := database.DB.Find(&todos).Error
	return todos, err
}

func (r *TodoRepository) GetByID(id uint) (*models.Todo, error) {
	var todo models.Todo
	err := database.DB.First(&todo, id).Error
	return &todo, err
}

func (r *TodoRepository) Update(todo *models.Todo) error {
	return database.DB.Save(todo).Error
}

func (r *TodoRepository) Delete(id uint) error {
	return database.DB.Delete(&models.Todo{}, id).Error
}
	