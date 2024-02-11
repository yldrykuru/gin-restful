package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CorsMiddleware is a Gin middleware function for handling Cross-Origin Resource Sharing (CORS).
func CorsMiddleware() gin.HandlerFunc {
	// Create a new CORS configuration with default settings.
	config := cors.DefaultConfig()

	// Allow requests from any origin. Limiting origins is recommended for security reasons.
	config.AllowOrigins = []string{"*"}

	// Specify allowed HTTP methods.
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

	// Specify allowed HTTP headers.
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}

	// Create and return a new CORS middleware instance with the configured settings.
	return cors.New(config)
}
