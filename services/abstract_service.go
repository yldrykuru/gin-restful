// services/abstract_service.go
package services

import (
	"yldrykuru/gin-restful/db"

	"gorm.io/gorm"
)

var database *gorm.DB = db.DB
