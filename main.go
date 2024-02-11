// main.go
package main

import (
	"fmt"
	"yldrykuru/gin-restful/configs"
	"yldrykuru/gin-restful/middlewares"
	"yldrykuru/gin-restful/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Yapılandırma ayarlarını yükle
	config := configs.GetConfig()

	// Servisleri oluşturun
	gin.SetMode(gin.DebugMode) // Gin'in çalışma modunu ayarlayın
	// Gin router'ını oluşturun
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	engine.SetTrustedProxies([]string{"192.168.1.1"})

	// Middleware'leri ekleyin
	engine.Use(middlewares.CorsMiddleware())
	engine.Use(middlewares.LoggerMiddleware())

	// Rotaları ayarlayın
	api := engine.Group("/api/v1")
	routes.SetupRoutes(api)

	// Server'ı başlatın
	port := config.Port
	if port == "" {
		port = "8080" // Varsayılan port
	}

	err := engine.Run(":" + port)
	if err != nil {
		fmt.Println("Server start failed:", err)
	}
}
