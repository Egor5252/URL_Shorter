package router

import (
	"url_shorter_new/internal/auth"
	"url_shorter_new/internal/handler"

	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	r.POST("/logout", handler.Logout)

	r.POST("/createshorturl", auth.AuthMiddleware(), handler.CreateShortUrl)
	r.POST("/go/:shorturl", handler.GoToShortUrl)
	r.POST("/account", auth.AuthMiddleware(), handler.Account)
	r.POST("/account/:id", auth.AuthMiddleware(), handler.UrlStatistics)
}
