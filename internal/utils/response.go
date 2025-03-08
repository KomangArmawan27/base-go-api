package utils

import (
	"github.com/gin-gonic/gin"
)

// BaseResponse formats error messages consistently
type BaseResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(c *gin.Context, code int, success bool, message string, data interface{}) {
	response := BaseResponse{
		Success: success,
		Code:    code,
		Message: message,
		Data:    data,
	}

	c.JSON(code, response)
}
