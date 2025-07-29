package handler

import (
	"net/http"
	"url_shorter_new/internal/auth"
	"url_shorter_new/internal/db"
	"url_shorter_new/internal/domain/user"
	"url_shorter_new/utils"

	"github.com/gin-gonic/gin"
)

var req struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func Register(c *gin.Context) {
	if err := c.Bind(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if req.Username == "" || req.Password == "" {
		utils.RespondError(c, http.StatusConflict, gin.H{
			"message": "Неверное заполнение полей формы регистрации",
		})
		return
	}

	val, err := db.ReadFirstByValue[user.User](user.UsersDB, "username", req.Username)
	if err != nil && err.Error() != "record not found" {
		utils.RespondError(c, http.StatusInternalServerError, gin.H{
			"message": "Ошибка поиска пользователя: " + err.Error(),
		})
		return
	}
	if val != nil {
		utils.RespondError(c, http.StatusConflict, gin.H{
			"message": "Логин занят",
		})
		return
	}

	passHash, err := utils.Hash(req.Password)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, gin.H{
			"message": "Ошибка хеширования пароля: " + err.Error(),
		})
		return
	}

	newUser := &user.User{
		Username: req.Username,
		PassHash: string(passHash),
	}

	if err := db.Create(user.UsersDB, newUser); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, gin.H{
			"message": "Ошибка создания пользователя в БД" + err.Error(),
		})
		return
	}

	err = auth.MakeJWT(c, newUser.ID, newUser.Username)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, gin.H{
			"message": "Ошибка создания JWT: " + err.Error() + ". Аккаунт создан, повторите попытку входа",
		})
		return
	}

	utils.RespondOK(c, gin.H{
		"message":  "Пользователь зарегистрирован, вход выполнен",
		"username": newUser.Username,
	})
}

func Login(c *gin.Context) {
	if err := c.Bind(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	findedUser, err := db.ReadFirstByValue[user.User](user.UsersDB, "username", req.Username)
	if err != nil {
		if err.Error() == "record not found" {
			utils.RespondError(c, http.StatusUnauthorized, gin.H{
				"message": "Ошибка логина или пароля",
			})
			return
		} else {
			utils.RespondError(c, http.StatusInternalServerError, gin.H{
				"message": "Ошибка поиска пользователя в БД: " + err.Error(),
			})
			return
		}
	}

	if err := utils.Compare(findedUser.PassHash, req.Password); err != nil {
		utils.RespondError(c, http.StatusUnauthorized, gin.H{
			"message": "Ошибка логина или пароля",
		})
		return
	}

	err = auth.MakeJWT(c, findedUser.ID, findedUser.Username)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, gin.H{
			"message": "Ошибка создания JWT: " + err.Error(),
		})
		return
	}

	utils.RespondOK(c, gin.H{
		"message":  "Вход выполнен",
		"username": findedUser.Username,
	})
}

func Logout(c *gin.Context) {
	auth.ResetCookie(c)
	utils.RespondOK(c, gin.H{
		"message": "Вы вышли из аккаунта",
	})
}
