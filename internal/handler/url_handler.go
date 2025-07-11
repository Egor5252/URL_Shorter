package handler

import (
	"fmt"
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

	html := `
		<!DOCTYPE html>
		<html lang="ru">
		<head>
			<meta charset="UTF-8">
			<title>Короткосыл</title>
			<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
			<style>
				body {
					background: #f8f9fa;
				}
				.centered-box {
					max-width: 400px;
					margin: 100px auto;
					padding: 30px;
					background: white;
					border-radius: 12px;
					box-shadow: 0 0 20px rgba(0, 0, 0, 0.1);
				}
			</style>
		</head>
		<body>
			<div class="centered-box">
				<h4 class="mb-4 text-center">Короткоссылка</h4>
				<form action="/createshorturl" method="POST">
					<div class="mb-3">
						<input type="text" class="form-control" name="url" value=` + fmt.Sprintf("192.168.1.182:8080/go/%s", newShortUrl.ShortCode) + ` placeholder="Введите ссылку..." required>
					</div>
				</form>
			</div>
		</body>
		</html>
	`

	// c.JSON(http.StatusOK, gin.H{"status": "короткая ссылка создана: " + newShortUrl.ShortCode})
	c.Data(200, "text/html; charset=utf-8", []byte(html))
}

func CreateShortUrlGet(c *gin.Context) {
	html := `
		<!DOCTYPE html>
		<html lang="ru">
		<head>
			<meta charset="UTF-8">
			<title>Короткосыл</title>
			<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
			<style>
				body {
					background: #f8f9fa;
				}
				.centered-box {
					max-width: 400px;
					margin: 100px auto;
					padding: 30px;
					background: white;
					border-radius: 12px;
					box-shadow: 0 0 20px rgba(0, 0, 0, 0.1);
				}
			</style>
		</head>
		<body>
			<div class="centered-box">
				<h4 class="mb-4 text-center">Введите ссылку</h4>
				<form action="/createshorturl" method="POST">
					<div class="mb-3">
						<input type="text" class="form-control" name="url" placeholder="Введите ссылку..." required>
					</div>
					<div class="d-grid">
						<button type="submit" class="btn btn-primary">Короткоссыльнуть</button>
					</div>
				</form>
			</div>
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
