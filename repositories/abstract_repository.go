// repositories/abstract_repository.go
package repositories

import (
	"yldrykuru/gin-restful/db"

	"gorm.io/gorm"
)

var database *gorm.DB = db.DB
