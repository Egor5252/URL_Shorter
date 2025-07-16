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

var jwtKey = []byte("jdd839jd73hksjfn332kfjng5ddu325jr")

func MakeJWT(id uint, name string) (string, error) {
	now := time.Now()
	expirationTime := now.Add(1 * time.Hour)

	claims := &Claims{
		ID:       id,
		Username: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Subject:   fmt.Sprintf("%d", id),
			Issuer:    "your-app-name",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("невозможно создать JWT токен: %w", err)
	}

	return tokenString, nil
}

func Who(c *gin.Context) (*Claims, error) {
	tokenStr, err := c.Cookie("token")
	if err != nil {
		return nil, fmt.Errorf("вход не выполнен")
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
