package routes

import (
	"yldrykuru/gin-restful/controllers"

	"github.com/gin-gonic/gin"
)

// todosRoutes configures routes related to todos within the specified RouterGroup.
func todosRoutes(engine *gin.RouterGroup) {
	// Create a new group for todos under the specified RouterGroup.
	todos := engine.Group("/todos")

	// Define routes for retrieving all todos and a specific todo by ID.
	todos.GET("/", controllers.GetTodos)
	todos.GET("/:id", controllers.GetTodo)

	// Define routes for adding, editing, and deleting todos.
	todos.POST("/", controllers.AddTodo)
	todos.PUT("/:id", controllers.EditTodo)
	todos.DELETE("/:id", controllers.DeleteTodo)
}
