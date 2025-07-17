package handler

import (
	"net/http"
	"time"
	"urlShorter/internal/auth"
	"urlShorter/internal/db"
	"urlShorter/internal/domain/url"
	"urlShorter/internal/domain/user"
	"urlShorter/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const secureCookie = false

func Register(c *gin.Context) {
	var req struct {
		User     string `json:"user" form:"user"`
		Password string `json:"password" form:"password"`
	}
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.User == "" || req.Password == "" {
		c.JSON(http.StatusConflict, gin.H{"error": "Неверное заполнение полей"})
		return
	}

	passHash, err := utils.Hash(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка хеша: " + err.Error()})
		return
	}

	val, err := db.ReadFirstByValue[user.User](user.UsersDB, "name", req.User)
	if err != nil && err.Error() != "record not found" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка поиска пользователя: " + err.Error()})
		return
	}
	if val != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Логин занят"})
		return
	}

	newUser := &user.User{
		Name:     req.User,
		PassHash: string(passHash),
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if err := db.Create(user.UsersDB, newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := auth.MakeJWT(newUser.ID, newUser.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания JWT: " + err.Error() + " . Аккаунт создан, повторите попытку входа"})
		return
	}

	c.SetCookie("token", tokenString, 60*60, "/", "", secureCookie, true)

	c.JSON(http.StatusCreated, gin.H{"message": gin.H{"info": "Пользователь зарегистрирован, вход выполнен", "user": newUser.Name}})
}

func Login(c *gin.Context) {
	var incomingUser struct {
		User     string `json:"user" form:"user"`
		Password string `json:"password" form:"password"`
	}

	if err := c.Bind(&incomingUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	findedUser, err := db.ReadFirstByValue[user.User](user.UsersDB, "name", incomingUser.User)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Ошибка логина или пароля"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка входа"})
			return
		}
	}

	if err := utils.Compare(findedUser.PassHash, incomingUser.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Ошибка логина или пароля"})
		return
	}

	tokenString, err := auth.MakeJWT(findedUser.ID, findedUser.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания JWT: " + err.Error()})
		return
	}

	c.SetCookie("token", tokenString, 60*60, "/", "", secureCookie, true)

	c.JSON(http.StatusOK, gin.H{"message": gin.H{"info": "Вход выполнен", "user": findedUser.Name}})
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", secureCookie, true)
	c.JSON(http.StatusOK, gin.H{"message": "Успешный выход"})
}

func Account(c *gin.Context) {
	claimsVal, ok := c.Get("claims")
	if !ok {
		c.JSON(500, gin.H{"error": "claims not found"})
		return
	}
	claims := claimsVal.(*auth.Claims)

	urls, err := db.ReadAllByValue[url.Url](url.UrlDB, "user_id", claims.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": urls})
}
