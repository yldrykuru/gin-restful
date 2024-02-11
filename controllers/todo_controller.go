package controllers

import (
	"net/http"
	"yldrykuru/gin-restful/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	// TODO: Get todos logic
	c.JSON(http.StatusOK, gin.H{"message": "GetTodos endpoint"})
}

func AddTodo(c *gin.Context) {
	var todo models.Todo

	// Gelen JSON verisini bind et
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Modeli doğrula
	if err := todo.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Todo'yu veritabanına ekle veya başka bir işlemi gerçekleştir

	c.JSON(http.StatusCreated, gin.H{"message": "Todo created successfully", "todo": todo})
}
func DeleteTodo(c *gin.Context) {
	// TODO: Delete todo logic
	c.JSON(http.StatusOK, gin.H{"message": "DeleteTodo endpoint"})
}
