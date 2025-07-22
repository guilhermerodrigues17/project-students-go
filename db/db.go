package db

import (
	"fmt"
	"log"

	"github.com/guilhermerodrigues17/project-students-go/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type StudentHandler struct {
	Db *gorm.DB
}

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=database port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(schemas.Student{})

	return db
}

func CreateStudentHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{Db: db}
}

func (s *StudentHandler) AddStudent(student schemas.Student) error {
	if result := s.Db.Create(&student); result.Error != nil {
		return result.Error
	}

	fmt.Println("Create student!")
	return nil
}

func (s *StudentHandler) GetStudents() ([]schemas.Student, error) {
	students := []schemas.Student{}

	err := s.Db.Find(&students).Error
	return students, err
}

func (s *StudentHandler) GetStudentsByActive(active bool) ([]schemas.Student, error) {
	filteredStudents := []schemas.Student{}

	err := s.Db.Where("active=?", active).Find(&filteredStudents).Error
	return filteredStudents, err
}

func (s *StudentHandler) GetStudent(id int) (schemas.Student, error) {
	student := schemas.Student{}

	err := s.Db.First(&student, id).Error
	return student, err
}

func (s *StudentHandler) UpdateStudent(updatedStudent schemas.Student) error {
	return s.Db.Save(&updatedStudent).Error
}

func (s *StudentHandler) DeleteStudent(student schemas.Student) error {
	return s.Db.Delete(&student).Error
}
