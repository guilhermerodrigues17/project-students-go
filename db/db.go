package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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
