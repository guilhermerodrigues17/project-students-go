package api

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermerodrigues17/project-students-go/db"
)

type Api struct {
	Gin *gin.Engine
	Db *db.StudentHandler
}

func NewServer() *Api {
	server := gin.Default()
	database := db.Init()
	studentDb := db.CreateStudentHandler(database)

	return &Api{
		Gin: server,
		Db: studentDb,
	}
}

func (api *Api) ConfigureRoutes() {
	api.Gin.GET("/ping", ping)
	api.Gin.GET("/students", api.getStudents)
	api.Gin.POST("/students", api.createStudent)
	api.Gin.GET("/students/:id", api.getStudent)
	api.Gin.PUT("/students/:id", api.updateStudent)
	api.Gin.DELETE("/students/:id", api.deleteStudent)
}

func (api *Api) Start(port ...string) error {
  	if len(port) > 0 {
		return api.Gin.Run(":" + port[0])	
	}
	return api.Gin.Run()
}
