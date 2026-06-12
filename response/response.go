package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {

	Success bool `json:"success"`

	Message string `json:"message,omitempty"`

	Data interface{} `json:"data,omitempty"`

	Pagination interface{} `json:"pagination,omitempty"`
}


func Success(
	c *gin.Context,
	data interface{},
) {

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data: data,
	})
}

func Created(
	c *gin.Context,
	data interface{},
) {

	c.JSON(http.StatusCreated, APIResponse{
		Success: true,
		Data: data,
	})
}

func Error(
	c *gin.Context,
	status int,
	message string,
) {

	c.JSON(status, APIResponse{
		Success: false,
		Message: message,
	})
}