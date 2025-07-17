package api

import (
	"fmt"
	"net/http"

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


func ping(c *gin.Context ) {
	c.JSON(http.StatusOK, gin.H{
	"message": "pong",
	})
}

func (api *Api) getStudents(c *gin.Context) {
	students, err := api.Db.GetStudents()
	if err != nil {
		c.String(http.StatusNotFound, "Failed to get students...")
	}

	c.JSON(http.StatusOK, students)
}

func (api *Api) createStudent(c *gin.Context) {
	student := db.Student{}
	
	if err := c.Bind(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := api.Db.AddStudent(student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})		
	}
	
	c.JSON(http.StatusCreated, "Create student")
}

func (api *Api) getStudent(c *gin.Context) {
	id := c.Param("id")
	printStr := fmt.Sprintf("Get %s user", id)
	c.String(http.StatusOK, printStr)
}

func (api *Api) updateStudent(c *gin.Context) {
	id := c.Param("id")
	printStr := fmt.Sprintf("Update %s user", id)
	c.String(http.StatusOK, printStr)
}

func (api *Api) deleteStudent(c *gin.Context) {
	id := c.Param("id")
	printStr := fmt.Sprintf("Delete %s user", id)
	c.String(http.StatusOK, printStr)
}
