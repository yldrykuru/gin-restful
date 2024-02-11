// services/todo_service.go
package services

import (
	"yldrykuru/gin-restful/models"
	"yldrykuru/gin-restful/repositories"
)

// TodoService defines the service layer for managing Todo entities.
type TodoService struct {
	// Additional fields can be added if needed.
}

// AddTodo defines the function to add a new Todo to the database.
func (s *TodoService) AddTodo(todo *models.Todo) (*models.Todo, error) {
	// Create a new Todo instance with the provided data.
	newTodo := models.Todo{
		Title:  todo.Title,
		Status: todo.Status,
	}

	// Add the new Todo to the database.
	database.Create(&newTodo)

	// No need to perform error checking since the Create function doesn't return an error.
	return &newTodo, nil
}

// EditTodo defines the function to edit an existing Todo in the database.
func (s *TodoService) EditTodo(todoID uint, todo *models.Todo) (*models.Todo, error) {
	// Retrieve the existing Todo from the database using its ID.
	todoRepository := repositories.TodoRepository{}
	existingTodo, err := todoRepository.GetTodoByID(todoID)
	if err != nil {
		return nil, err
	}

	// Update the existing Todo with the provided data.
	existingTodo.Title = todo.Title
	existingTodo.Status = todo.Status

	// Save the updated Todo to the database.
	database.Save(existingTodo)

	return existingTodo, nil
}

// DeleteTodo defines the function to delete an existing Todo from the database.
func (s *TodoService) DeleteTodo(todoID uint) error {
	// Retrieve the existing Todo from the database using its ID.
	todoRepository := repositories.TodoRepository{}
	existingTodo, err := todoRepository.GetTodoByID(todoID)
	if err != nil {
		return err
	}

	// Delete the Todo from the database.
	database.Delete(existingTodo)

	return nil
}
