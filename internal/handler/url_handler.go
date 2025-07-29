package handler

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"url_shorter_new/internal/auth"
	"url_shorter_new/internal/db"
	"url_shorter_new/internal/domain/url"
	"url_shorter_new/internal/domain/visits"
	"url_shorter_new/utils"

	"github.com/gin-gonic/gin"
)

func CreateShortUrl(c *gin.Context) {
	claims, err := auth.GetClaims(c)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var req struct {
		Url       string `json:"url" form:"url"`
		ShortCode string `json:"short_code" form:"short_code"`
	}

	if err := c.Bind(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if !strings.HasPrefix(req.Url, "http://") && !strings.HasPrefix(req.Url, "https://") {
		req.Url = "https://" + req.Url
	}

	if req.ShortCode == "" {
		word, err := utils.RandomWord()
		if err != nil {
			utils.RespondError(c, http.StatusInternalServerError, gin.H{
				"message": "Ошибка получения случайного слова: " + err.Error(),
			})
			return
		}

		req.ShortCode = fmt.Sprintf("%s%d", word, time.Now().Unix())
	}

	newShortUrl := &url.Url{
		UserID:      claims.ID,
		OriginalURL: req.Url,
		ShortCode:   req.ShortCode,
	}

	if err := db.Create(url.UrlDB, newShortUrl); err != nil {
		if err.Error() == "constraint failed: UNIQUE constraint failed: urls.short_code (2067)" {
			utils.RespondError(c, http.StatusConflict, gin.H{
				"message": "Ссылка уже занята",
			})
		} else {
			utils.RespondError(c, http.StatusInternalServerError, gin.H{
				"message": "Ошибка добавления ссылки в БД: " + err.Error(),
			})
		}
		return
	}

	utils.RespondOK(c, gin.H{
		"message": fmt.Sprintf("localhost:8080/go/%s", newShortUrl.ShortCode),
	})
}

func GoToShortUrl(c *gin.Context) {
	shortUrl := c.Param("shorturl")

	findedUrl, err := db.ReadFirstByValue[url.Url](url.UrlDB, "short_code", shortUrl)
	if err != nil {
		if err.Error() == "record not found" {
			utils.RespondError(c, http.StatusBadGateway, gin.H{
				"message": "Несуществующая ссылка",
			})
		} else {
			utils.RespondError(c, http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("Ошибка поиска в БД ссылок: %s", err.Error()),
			})
		}
		return
	}

	newVisit := &visits.Visits{
		ShortUrlID: findedUrl.ID,
		UserIP:     c.ClientIP(),
	}

	err = db.Create(visits.VisitsDB, newVisit)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Ошибка создания в БД визитов: %s", err.Error()),
		})
		return
	}

	c.Redirect(http.StatusFound, findedUrl.OriginalURL)
}
