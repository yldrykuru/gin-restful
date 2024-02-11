package routes

import (
	"yldrykuru/gin-restful/controllers"

	"github.com/gin-gonic/gin"
)

func todosRoutes(engine *gin.RouterGroup) {
	todos := engine.Group("/todos")
	todos.GET("/", controllers.GetTodos)
	todos.POST("/", controllers.AddTodo)
	todos.DELETE("/:id", controllers.DeleteTodo)
}
