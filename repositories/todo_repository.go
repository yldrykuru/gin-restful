// repositories/todo_repository.go
package repositories

import (
	"yldrykuru/gin-restful/models"
)

// TodoRepository defines the repository for managing Todo models in the database.
type TodoRepository struct{}

// GetTodos retrieves a list of todos from the database with pagination.
// It takes 'page' and 'pageSize' parameters to specify the page number and
// the number of todos per page. The result is a slice of todos and an error.
// If successful, the function returns the requested todos, otherwise an error.
func (r *TodoRepository) GetTodos(page, pageSize int) ([]models.Todo, error) {
	// Slice to store the retrieved todos
	var todos []models.Todo

	// Perform database query with pagination using Offset and Limit
	result := database.Offset((page - 1) * pageSize).Limit(pageSize).Find(&todos)
	if result.Error != nil {
		// Return an error if the database query encounters an issue
		return nil, result.Error
	}

	// Return the retrieved todos and no error in case of success
	return todos, nil
}

// GetTodoByID retrieves a todo from the database based on its ID.
func (r *TodoRepository) GetTodoByID(id uint) (*models.Todo, error) {
	var todo models.Todo

	// Find a todo by its ID from the database.
	result := database.First(&todo, id)
	if result.Error != nil {
		// Return an error if there is an issue fetching the todo.
		return nil, result.Error
	}

	// Return the retrieved todo.
	return &todo, nil
}
