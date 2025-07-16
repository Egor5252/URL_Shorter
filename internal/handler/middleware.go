package handler

import (
	"net/http"
	"urlShorter/internal/auth"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := auth.Who(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
