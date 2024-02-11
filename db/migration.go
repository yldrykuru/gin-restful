// db/migration.go
package db

import (
	"yldrykuru/gin-restful/models"
)

// AutoMigrate function performs automatic database migration for defined models.
func AutoMigrate() {
	// AutoMigrate function automatically creates or updates database tables
	// based on the provided model structs.
	DB.AutoMigrate(&models.Todo{})
	// If you want to add more models, you can include them here.
}
