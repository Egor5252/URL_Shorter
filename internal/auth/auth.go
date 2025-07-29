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

	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	access_tokenString, err := access_token.SignedString(accessJwtKey)
	if err != nil {
		return fmt.Errorf("невозможно создать JWT токен: %w", err)
	}

	c.SetCookie("access_token", access_tokenString, CookieLiveTime, "/", "", secureCookie, true)

	return nil
}

func Who(c *gin.Context) (*Claims, error) {
	access_tokenStr, err := c.Cookie("access_token")
	if err != nil || access_tokenStr == "" {
		return nil, fmt.Errorf("нет токена")
	}

	claims := &Claims{}
	access_token, err := jwt.ParseWithClaims(access_tokenStr, claims, func(access_token *jwt.Token) (any, error) {
		if _, ok := access_token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неожиданный метод подписи: %v", access_token.Header["alg"])
		}
		return accessJwtKey, nil
	})

	if err != nil || !access_token.Valid {
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
	c.SetCookie("access_token", "", -1, "/", "", secureCookie, true)
}
