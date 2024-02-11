// controllers/todo_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"yldrykuru/gin-restful/models"
	"yldrykuru/gin-restful/repositories"
	"yldrykuru/gin-restful/services"

	"github.com/gin-gonic/gin"
)

// GetTodos is a controller function to handle the retrieval of paginated todos.
func GetTodos(c *gin.Context) {
	// Initialize TodoRepository
	todoRepository := repositories.TodoRepository{}

	// Retrieve page and pageSize from query parameters with default values
	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	if pageSize <= 0 {
		pageSize = 10 // Default page size
	}

	// Retrieve paginated todos from the repository
	todos, err := todoRepository.GetTodos(page, pageSize)
	if err != nil {
		// If there is an error fetching todos, respond with an Internal Server Error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}

	// Respond with the retrieved paginated todos
	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

// GetTodo function to retrieve a specific todo by ID
func GetTodo(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	todoRepository := repositories.TodoRepository{}
	todo, err := todoRepository.GetTodoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

// AddTodo function to add a new todo
func AddTodo(c *gin.Context) {
	var todo models.Todo

	// Bind JSON data
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the model
	if err := todo.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add todo to the database
	todoService := services.TodoService{}
	data, err := todoService.AddTodo(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add todo to the database"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Todo created successfully", "todo": data})
}

// EditTodo function to edit an existing todo
func EditTodo(c *gin.Context) {
	// Get todo ID from the request
	todoID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	var updatedTodo models.Todo

	// Bind JSON data
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the model
	if err := updatedTodo.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Edit the todo in the database
	todoService := services.TodoService{}
	data, err := todoService.EditTodo(uint(todoID), &updatedTodo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit todo in the database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully", "todo": data})
}

// DeleteTodo function to delete an existing todo
func DeleteTodo(c *gin.Context) {
	// Get todo ID from the request
	todoID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	// Delete the todo from the database
	todoService := services.TodoService{}
	err = todoService.DeleteTodo(uint(todoID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo from the database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
