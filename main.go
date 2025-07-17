package main

import (
	"log"

	"github.com/guilhermerodrigues17/project-students-go/api"
)

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
