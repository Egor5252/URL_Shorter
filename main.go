package main

import (
	"fmt"
	transitionstatistics "urlShorter/internal/domain/transitionStatistics"
	"urlShorter/internal/domain/url"
	"urlShorter/internal/domain/user"
	"urlShorter/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	user.InitDB()
	url.InitDB()
	transitionstatistics.InitDB()

	r := gin.Default()
	router.InitRouters(r)

	if err := r.Run(":8080"); err != nil {
		fmt.Println(err.Error())
	}
}
