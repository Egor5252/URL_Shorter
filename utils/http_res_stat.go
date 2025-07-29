package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ответ с данными
func RespondOK(c *gin.Context, data gin.H) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
		"error":   nil,
	})
}

// Ответ с ошибкой
func RespondError(c *gin.Context, httpStatus int, err gin.H) {
	c.JSON(httpStatus, gin.H{
		"success": false,
		"data":    nil,
		"error":   err,
	})
}
