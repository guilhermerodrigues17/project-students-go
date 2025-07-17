package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type StudentHandler struct{
	Db *gorm.DB
}

type Student struct {
	gorm.Model
	Name string `json:"name"`
	Cpf string `json:"cpf"`
	Email string `json:"email"`
	Age int `json:"age"`
	Active bool `json:"active"`
}

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=database port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(Student{})

	return db
}

func CreateStudentHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{Db: db}
}


func (s *StudentHandler) AddStudent(student Student) error {
	if result := s.Db.Create(&student); result.Error != nil {
		return result.Error
	}

	fmt.Println("Create student!")
	return nil
}

func (s *StudentHandler) GetStudents() ([]Student, error) {
	students := []Student{}

	err := s.Db.Find(&students).Error
	return students, err
}

func (s *StudentHandler) GetStudent(id int) (Student, error) {
	student := Student{}

	err := s.Db.First(&student, id).Error
	return student, err	
}