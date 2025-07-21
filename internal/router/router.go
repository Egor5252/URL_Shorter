package router

import (
	"url_shorter_new/internal/auth"
	"url_shorter_new/internal/handler"

	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	r.Static("assets", "web-app/dist/assets")
	// r.StaticFile("vite.svg", "web-app/dist/vite.svg")
	r.LoadHTMLFiles("web-app/dist/index.html")
	r.GET("/", handler.RootHandler)

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	r.POST("/logout", handler.Logout)

	r.POST("/createshorturl", auth.AuthMiddleware(), handler.CreateShortUrl)
	r.GET("/go/:shorturl", handler.GoToShortUrl)
	r.GET("/account", auth.AuthMiddleware(), handler.Account)
	r.GET("/account/:id", auth.AuthMiddleware(), handler.UrlStatistics)
}
