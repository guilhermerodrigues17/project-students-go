package api

import "github.com/gin-gonic/gin"

type httpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewErr(c *gin.Context, status int, err error) {
	er := httpError{
		Code:    status,
		Message: err.Error(),
	}

	c.JSON(status, er)
}
