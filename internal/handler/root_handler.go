package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func NoPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
