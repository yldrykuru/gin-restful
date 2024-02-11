// db/db.go
package db

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is the global variable for the Gorm database connection.
var DB *gorm.DB

// init function sets up the database connection when the package is initialized.
func init() {
	// Create a DSN (Data Source Name) for the MySQL connection.
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") +
		"@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" +
		os.Getenv("DB_NAME") + "?parseTime=true"

	// Open a Gorm database connection.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database connection failed: " + err.Error())
	}

	// Set the global DB variable.
	DB = db
}
