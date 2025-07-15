package handler

import (
	"net/http"
	"urlShorter/internal/auth"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	claims, err := auth.Who(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.Set("claims", claims)
	c.Next()
}
