package handler

import (
	"net/http"
	"strings"
	"time"
	"urlShorter/internal/auth"
	"urlShorter/internal/db"
	transitionstatistics "urlShorter/internal/domain/transitionStatistics"
	"urlShorter/internal/domain/url"
	"urlShorter/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateShortUrl(c *gin.Context) {
	claimsVal, ok := c.Get("claims")
	if !ok {
		c.JSON(500, gin.H{"error": "claims not found"})
		return
	}
	claims := claimsVal.(*auth.Claims)

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
		UserID:      claims.ID,
		OriginalURL: req.Url,
		ShortCode:   shortUrl,
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if err := db.Create(url.UrlDB, newShortUrl); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{"message": "10.10.13.40:8080/go/" + newShortUrl.ShortCode},
	)
}

func GoToShortUrl(c *gin.Context) {
	shortUrl := c.Param("shorturl")

	findedUrl, err := db.ReadFirstByValue[url.Url](url.UrlDB, "short_code", shortUrl)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": err.Error()})
		return
	}

	newVal := &transitionstatistics.Transitionstatistics{
		UserIP:   c.ClientIP(),
		ShortUrl: shortUrl,
		Model: gorm.Model{
			CreatedAt: time.Now(),
		},
	}

	err = db.Create(transitionstatistics.TSDB, newVal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, findedUrl.OriginalURL)
}
