package main

import (
	"fmt"
	"url_shorter_new/internal/domain/url"
	"url_shorter_new/internal/domain/user"
	"url_shorter_new/internal/domain/visits"
	"url_shorter_new/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	user.InitDB()
	url.InitDB()
	visits.InitDB()

	r := gin.Default()
	router.InitRouters(r)

	if err := r.Run(":8080"); err != nil {
		fmt.Println(err.Error())
	}
}
