// utils/response.go
package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func JSONSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, SuccessResponse{Message: message, Data: data})
}

func JSONError(c *gin.Context, statusCode int, errMessage string) {
	c.JSON(statusCode, ErrorResponse{Error: errMessage})
}
