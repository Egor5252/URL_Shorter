package main

import (
	"fmt"
	"url_shorter_new/internal/domain/login"
	"url_shorter_new/internal/domain/url"
	"url_shorter_new/internal/domain/user"
	"url_shorter_new/internal/domain/visits"
	"url_shorter_new/internal/router"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	user.InitDB()
	url.InitDB()
	visits.InitDB()
	login.InitDB()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true, // ВАЖНО
	}))
	router.InitRouters(r)

	if err := r.Run(":8080"); err != nil {
		fmt.Println(err.Error())
	}
}
