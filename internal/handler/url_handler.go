package handler

import (
	"net/http"
	"strings"
	"time"
	"urlShorter/internal/db"
	"urlShorter/internal/domain/url"
	"urlShorter/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateShortUrl(c *gin.Context) {
	var req struct {
		Url string `json:"url" form:"url"`
	}

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !strings.HasPrefix(req.Url, "http://") && !strings.HasPrefix(req.Url, "https://") {
		req.Url = "https://" + req.Url
	}

	shortUrl, err := utils.RandomWord()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newShortUrl := &url.Url{
		OriginalURL: req.Url,
		ShortCode:   string(shortUrl),
		Count:       0,
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if err := db.Create(url.UrlDB, newShortUrl); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "короткая ссылка создана: " + newShortUrl.ShortCode})
}

func CreateShortUrlGet(c *gin.Context) {
	html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Форма с полем ввода</title>
		</head>
		<body>
			<h1>Введите ссылку и нажмите кнопку</h1>
			<form action="/createshorturl" method="POST">
				<input type="text" name="url" placeholder="Введите ссылку" required>
				<button type="submit">Отправить</button>
			</form>
		</body>
		</html>
	`
	c.Data(200, "text/html; charset=utf-8", []byte(html))
}

func GoToShortUrl(c *gin.Context) {
	shortUrl := c.Param("shorturl")

	findedUrl, err := db.ReadFirstByValue[url.Url](url.UrlDB, "short_code", shortUrl)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": err.Error()})
		return
	}

	db.Update(url.UrlDB, findedUrl, "count", findedUrl.Count+1)

	c.Redirect(http.StatusFound, findedUrl.OriginalURL)
}
