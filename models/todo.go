// validators/todo.go
package models

import "github.com/go-playground/validator/v10"

// Todo modelini tanımla
type Todo struct {
	ID     int    `json:"id" validate:"required"`
	Title  string `json:"title" validate:"required,min=3,max=50"`
	Status bool   `json:"status" validate:"required"`
}

// Validate fonksiyonunu tanımla
func (t *Todo) Validate() error {
	validate := validator.New()

	// Struct tag'lerine göre doğrulama yap
	return validate.Struct(t)
}
