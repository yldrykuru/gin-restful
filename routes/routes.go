package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.RouterGroup) {
	todosRoutes(engine)
}
