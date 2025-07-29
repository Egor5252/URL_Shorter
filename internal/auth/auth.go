package auth

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWT
type Claims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Время жизни куки в cекундах
const CookieLiveTime = 60 * 60
const secureCookie = false

var accessJwtKey = []byte("jdd839jd73hksjfn332kfjng5ddu325jr322")

func MakeJWT(c *gin.Context, id uint, name string) error {
	now := time.Now()
	expirationTime := now.Add(CookieLiveTime * time.Second)

	claims := &Claims{
		ID:       id,
		Username: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Subject:   fmt.Sprintf("%d", id),
			Issuer:    "korotkosill",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(accessJwtKey)
	if err != nil {
		return fmt.Errorf("невозможно создать JWT токен: %w", err)
	}

	c.SetCookie("token", tokenString, CookieLiveTime, "/", "", secureCookie, true)

	return nil
}

func Who(c *gin.Context) (*Claims, error) {
	tokenStr, err := c.Cookie("token")
	if err != nil || tokenStr == "" {
		return nil, fmt.Errorf("нет токена")
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неожиданный метод подписи: %v", token.Header["alg"])
		}
		return accessJwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("недействительный или просроченный токен")
	}

	return claims, nil

}

func GetClaims(c *gin.Context) (*Claims, error) {
	claimsVal, ok := c.Get("claims")
	if !ok {
		return nil, fmt.Errorf("структура Claims не найдена")
	}

	return claimsVal.(*Claims), nil
}

func ResetCookie(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", secureCookie, true)
}
