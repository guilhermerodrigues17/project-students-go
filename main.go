package main

import (
	"fmt"
	"net/http"

	"github.com/guilhermerodrigues17/project-students-go/db"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
  	
	router.GET("/ping", ping)
  	router.GET("/students", getStudents)
	router.POST("/students", createStudent)
	router.GET("/students/:id", getStudent)
	router.PUT("/students/:id", updateStudent)
	router.DELETE("/students/:id", deleteStudent)

  	router.Run() //listen and serve on localhost:8080
}

func ping(c *gin.Context ) {
	c.JSON(http.StatusOK, gin.H{
	"message": "pong",
	})
}

func getStudents(c *gin.Context) {
	c.String(http.StatusOK, "List of all students")
}

func createStudent(c *gin.Context) {
	student := db.Student{}
	
	if err := c.Bind(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := db.AddStudent(student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})		
	}
	
	c.JSON(http.StatusCreated, "Create student")
}

func getStudent(c *gin.Context) {
	id := c.Param("id")
	printStr := fmt.Sprintf("Get %s user", id)
	c.String(http.StatusOK, printStr)
}

func updateStudent(c *gin.Context) {
	id := c.Param("id")
	printStr := fmt.Sprintf("Update %s user", id)
	c.String(http.StatusOK, printStr)
}

func deleteStudent(c *gin.Context) {
	id := c.Param("id")
	printStr := fmt.Sprintf("Delete %s user", id)
	c.String(http.StatusOK, printStr)
}



