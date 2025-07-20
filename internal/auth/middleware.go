package auth

import (
	"net/http"
	"url_shorter_new/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := Who(c)
		if err != nil {
			utils.RespondError(c, http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
