// models/todo.go
package models

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel represents the basic structure for database models.
type BaseModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
