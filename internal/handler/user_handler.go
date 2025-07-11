package handler

import (
	"net/http"
	"time"
	"urlShorter/internal/auth"
	"urlShorter/internal/db"
	"urlShorter/internal/domain/user"
	"urlShorter/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const secureCookie = false

func Register(c *gin.Context) {
	var req struct {
		Name     string `json:"name" form:"name"`
		Password string `json:"password" form:"password"`
	}
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passHash, err := utils.Hash(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка хеша: " + err.Error()})
		return
	}

	val, err := db.ReadFirstByValue[user.User](user.UsersDB, "name", req.Name)
	if err != nil && err.Error() != "record not found" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка поиска пользователя: " + err.Error()})
		return
	}
	if val != nil {
		c.JSON(http.StatusConflict, gin.H{"status": "Логин занят"})
		return
	}

	newUser := &user.User{
		Name:     req.Name,
		PassHash: string(passHash),
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if err := db.Create(user.UsersDB, newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	tokenString, err := auth.MakeJWT(newUser.ID, newUser.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Ошибка создания JWT: " + err.Error()})
		return
	}

	c.SetCookie("token", tokenString, 60*60, "/", "", secureCookie, true)

	c.JSON(http.StatusCreated, gin.H{"status": gin.H{"info": "Пользователь зарегистрирован, вход выполнен", "user": newUser.Name}})
}

func Login(c *gin.Context) {
	var incomingUser struct {
		Name     string `json:"name" form:"name"`
		Password string `json:"password" form:"password"`
	}

	if err := c.Bind(&incomingUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	findedUser, err := db.ReadFirstByValue[user.User](user.UsersDB, "name", incomingUser.Name)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "Ошибка логина или пароля"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "Ошибка входа"})
			return
		}
	}

	if err := utils.Compare(findedUser.PassHash, incomingUser.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "Ошибка логина или пароля"})
		return
	}

	tokenString, err := auth.MakeJWT(findedUser.ID, findedUser.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Ошибка создания JWT: " + err.Error()})
		return
	}

	c.SetCookie("token", tokenString, 60*60, "/", "", secureCookie, true)

	c.JSON(http.StatusOK, gin.H{"status": gin.H{"info": "Вход выполнен", "user": findedUser.Name}})
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", secureCookie, true)
	c.JSON(http.StatusOK, gin.H{"message": "Успешный выход"})
}

func Account(c *gin.Context) {
	incomingUser, err := auth.Who(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": incomingUser})
}
