package router

import (
	"urlShorter/internal/handler"

	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	r.POST("/logout", handler.Logout)
	r.GET("/account", handler.Account)

	r.POST("/createshorturl", handler.AuthMiddleware, handler.CreateShortUrl)
	r.GET("/go/:shorturl", handler.GoToShortUrl)
}
