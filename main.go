// main.go
package main

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"os"
	"yldrykuru/gin-restful/db"
	"yldrykuru/gin-restful/middlewares"
	"yldrykuru/gin-restful/routes"
)

func main() {
	// Set the Gin operating mode based on the ENV value
	setGinMode()

	// Create a Gin router
	engine := setupRouter()

	// Perform database migration
	db.AutoMigrate()

	// Start the server
	port := os.Getenv("GIN_PORT") // GIN_PORT can be used instead of PORT
	if port == "" {
		port = "8080" // Default port
	}

	// Start the server on the specified port
	err := engine.Run(":" + port)
	if err != nil {
		// Log a fatal error if the server fails to start
		log.Fatalf("Failed to start the server: %s", err)
	}

}

// setupRouter configures the Gin router and sets up middleware.
func setupRouter() *gin.Engine {
	// Create a new Gin engine
	engine := gin.New()
	engine.SetTrustedProxies([]string{"127.0.0.1"})
	// Add middleware to the engine
	engine.Use(gin.Logger(), gin.Recovery())
	engine.Use(middlewares.CorsMiddleware())

	// Convert RATE_LIMIT to an integer
	rateLimit, err := strconv.Atoi(os.Getenv("RATE_LIMIT"))
	if err != nil {
		rateLimit = 100 // Default rate limit
	}

	engine.Use(middlewares.RateLimitMiddleware(rateLimit, time.Second))

	// Configure routes
	api := engine.Group("/api/v1")
	routes.SetupRoutes(api)

	return engine
}

// setGinMode sets the Gin operating mode based on the ENV value.
func setGinMode() {
	// Check the value obtained from ENV
	ginMode := os.Getenv("GIN_MODE")
	mode := gin.ReleaseMode
	if ginMode == "" || ginMode == "debug" {
		// Use DebugMode as the default if not specified in ENV
		mode = gin.DebugMode
	}

	// Set the Gin operating mode
	gin.SetMode(mode)
}
