package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilhermerodrigues17/project-students-go/db"
)

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
