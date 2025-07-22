package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string `json:"name"`
	Cpf    string `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active *bool  `json:"active"`
}

type StudentResponse struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	Active    *bool     `json:"active"`
}

func FormatedResponse(students []Student) []StudentResponse {
	var response []StudentResponse
	for _, student := range students {
		response = append(response, StudentResponse{
			ID:        int(student.ID),
			CreatedAt: student.CreatedAt,
			UpdatedAt: student.UpdatedAt,
			Name:      student.Name,
			Cpf:       student.Cpf,
			Email:     student.Email,
			Age:       student.Age,
			Active:    student.Active,
		})
	}
	return response
}
