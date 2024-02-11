package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures various route groups within the specified RouterGroup.
func SetupRoutes(engine *gin.RouterGroup) {
	// Configure routes related to todos.
	todosRoutes(engine)

	// Configure authentication-related routes.
	authRoutes(engine)
}
