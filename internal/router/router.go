package router

import (
	"urlShorter/internal/auth"
	"urlShorter/internal/handler"

	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	r.Static("assets", "web-app/dist/assets")
	r.LoadHTMLFiles("web-app/dist/index.html")

	r.GET("/", handler.RootHandler)

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	r.POST("/logout", handler.Logout)
	r.GET("/account", auth.AuthMiddleware(), handler.Account)

	r.POST("/createshorturl", auth.AuthMiddleware(), handler.CreateShortUrl)
	r.GET("/go/:shorturl", handler.GoToShortUrl)
}
