// db/db.go
package db

import (
	"yldrykuru/gin-restful/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	config := configs.GetConfig()

	// MySQL bağlantısı kurun
	dsn := config.DatabaseURL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database connection failed: " + err.Error())
	}

	// DB değişkenini ayarlayın
	DB = db
}
