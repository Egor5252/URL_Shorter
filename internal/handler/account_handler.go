package handler

import (
	"net/http"
	"strconv"
	"url_shorter_new/internal/auth"
	"url_shorter_new/internal/db"
	"url_shorter_new/internal/domain/url"
	"url_shorter_new/internal/domain/visits"
	"url_shorter_new/utils"

	"github.com/gin-gonic/gin"
)

func Account(c *gin.Context) {
	claims, err := auth.GetClaims(c)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	urls, err := db.ReadAllByValue[url.Url](url.UrlDB, "user_id", claims.ID)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	utils.RespondOK(c, gin.H{
		"message": urls,
	})
}

func UrlStatistics(c *gin.Context) {
	claims, err := auth.GetClaims(c)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	getID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Проверка принадлежности ссылки пользователю
	_, err = db.ReadOneByValues[url.Url](url.UrlDB, map[string]any{
		"id":      getID,
		"user_id": claims.ID,
	})
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, gin.H{
			"message": "Неизвестная ссылка",
		})
		return
	}

	// Получение статистики переходов
	visits, err := db.ReadAllByValue[visits.Visits](visits.VisitsDB, "short_url_id", getID)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	utils.RespondOK(c, gin.H{
		"message": visits,
	})
}
