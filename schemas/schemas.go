package schemas

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string `json:"name"`
	Cpf string `json:"cpf"`
	Email string `json:"email"`
	Age int `json:"age"`
	Active *bool `json:"active"`
}
