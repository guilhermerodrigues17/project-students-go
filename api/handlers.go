package api

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guilhermerodrigues17/project-students-go/schemas"
	"gorm.io/gorm"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary Ping example
// @Schemes
// @Description Faz uma requisição para health check da API
// @Tags Ping
// @Accept json
// @Produce json
// @Success 200 {string} pong
// @Router /ping [get]
func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func (api *Api) getStudents(c *gin.Context) {
	students, err := api.Db.GetStudents()
	if err != nil {
		NewErr(c, http.StatusInternalServerError, err)
		return
	}

	activeParam := c.Query("active")
	if activeParam != "" {
		active, err := strconv.ParseBool(activeParam)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		students, err = api.Db.GetStudentsByActive(active)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	listOfStudents := map[string][]schemas.StudentResponse{"students": schemas.FormatedResponse(students)}

	c.JSON(http.StatusOK, listOfStudents)
}

func (api *Api) createStudent(c *gin.Context) {
	studentReq := StudentRequest{}

	if err := c.Bind(&studentReq); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := studentReq.Validate(); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student := schemas.Student{
		Name:   studentReq.Name,
		Cpf:    studentReq.Cpf,
		Email:  studentReq.Email,
		Age:    studentReq.Age,
		Active: studentReq.Active,
	}

	if err := api.Db.AddStudent(student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, student)
}

func (api *Api) getStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusInternalServerError, "An error occurred...")
		return
	}

	student, err := api.Db.GetStudent(id)
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

	updatingStudent, err := api.Db.GetStudent(id)
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

	student, err := api.Db.GetStudent(id)
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
