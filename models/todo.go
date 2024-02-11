// validators/todo.go
package models

import (
	"github.com/go-playground/validator/v10"
)

// Todo represents the model for a todo item.
type Todo struct {
	BaseModel
	Title  string `json:"title" validate:"required,min=3,max=50"`
	Status bool   `json:"status" validate:"required"`
}

// Validate performs validation on the Todo model using struct tags.
func (t *Todo) Validate() error {
	validate := validator.New()

	// Perform validation based on struct tags.
	return validate.Struct(t)
}
