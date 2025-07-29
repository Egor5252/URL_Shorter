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

var jwtKey = []byte("jdd839jd73hksjfn332kfjng5ddu325jr322")

func MakeJWT(c *gin.Context, id uint, name string) (string, error) {
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
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("невозможно создать JWT токен: %w", err)
	}

	c.SetCookie("token", tokenString, CookieLiveTime, "/", "", secureCookie, true)

	return tokenString, nil
}

func Who(c *gin.Context) (*Claims, error) {
	tokenStr, err := c.Cookie("token")
	if err != nil || tokenStr == "" {
		// Если нет токена в cookie, ищем в Authorization Header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			return nil, fmt.Errorf("вход не выполнен, токен не найден")
		}

		const prefix = "Bearer "
		if len(authHeader) <= len(prefix) || authHeader[:len(prefix)] != prefix {
			return nil, fmt.Errorf("некорректный формат Authorization header")
		}

		tokenStr = authHeader[len(prefix):]
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неожиданный метод подписи: %v", token.Header["alg"])
		}
		return jwtKey, nil
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
