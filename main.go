package main

import (
	"log"

	"github.com/guilhermerodrigues17/project-students-go/api"
	_ "github.com/guilhermerodrigues17/project-students-go/docs"
)

// @title Students API
// @version 1.0
// @description Esta é uma API de Students com Gin + Swagger
// @host localhost:8080
// @BasePath /
func main() {
	//New instance of Gin engine
	server := api.NewServer()

	//Configure endpoints
	server.ConfigureRoutes()

	//Listen and serve on localhost:8080 by default
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
