package routes

import (
	"yldrykuru/gin-restful/controllers"
	"yldrykuru/gin-restful/middlewares"

	"github.com/gin-gonic/gin"
)

// authRoutes defines authentication-related routes under a given RouterGroup.
func authRoutes(engine *gin.RouterGroup) {
	// Create a new RouterGroup for authentication routes.
	auth := engine.Group("/auth")

	// Apply the AuthMiddleware to all routes under the /auth group.
	auth.Use(middlewares.AuthMiddleware())

	// Define a route to handle GET requests at /auth/.
	// This route is protected by the AuthMiddleware.
	auth.GET("/", controllers.GetTodos)
}
