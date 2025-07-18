package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guilhermerodrigues17/project-students-go/schemas"
	"gorm.io/gorm"
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
	student := schemas.Student{}
	
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusInternalServerError, "An error occurred...")
		return
	}

	student, err  := api.Db.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.String(http.StatusNotFound, "Student not found")
		return
	}

	if err != nil {
		c.String(http.StatusInternalServerError, "An error occurred...")
		return
	}

	c.JSON(http.StatusOK, student)
}

func (api *Api) updateStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusInternalServerError, "An error occurred...")
		return
	}

	updatingStudent, err  := api.Db.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.String(http.StatusNotFound, "Student not found")
		return
	}

	receivedStudent := schemas.Student{}
	if err := c.Bind(&receivedStudent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updatedStudent := verifyUpdateFields(receivedStudent, updatingStudent)

	if err := api.Db.UpdateStudent(updatedStudent); err != nil {
		c.String(http.StatusInternalServerError, "An error occurred...")
		return
	}

	c.JSON(http.StatusOK, updatedStudent)
}

func (api *Api) deleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusInternalServerError, "An error occurred...")
		return
	}

	student, err  := api.Db.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.String(http.StatusNotFound, "Student not found")
		return
	}

	if err := api.Db.DeleteStudent(student); err != nil {
		c.String(http.StatusInternalServerError, "An error occurred...")
		return
	}

	c.JSON(http.StatusOK, student)
}

func verifyUpdateFields(receivedStudent, updatingStudent schemas.Student) schemas.Student {

	if receivedStudent.Name != "" {
		updatingStudent.Name = receivedStudent.Name
	}

	if receivedStudent.Cpf != "" {
		updatingStudent.Cpf = receivedStudent.Cpf
	}

	if receivedStudent.Email != "" {
		updatingStudent.Email = receivedStudent.Email
	}

	if receivedStudent.Age > 0 {
		updatingStudent.Age = receivedStudent.Age
	}

	if receivedStudent.Active != nil {
		updatingStudent.Active = receivedStudent.Active
	}

	return updatingStudent
}
