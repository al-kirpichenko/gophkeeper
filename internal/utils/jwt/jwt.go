package jwt

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/al-kirpichenko/gophkeeper/internal/models"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID uint
	Login  string
}

const SecretKey = "bvaEFBtr5e"

// NewToken creates new JWT token for given user and app.
func NewToken(user *models.User, duration time.Duration) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
		UserID: user.ID,
		Login:  user.Login,
	})

	// Подписываем токен, используя секретный ключ приложения
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

//func ValidationToken(tokenString string) bool {
//
//	claims := &Claims{}
//	token, err := jwt.ParseWithClaims(tokenString, claims,
//		func(t *jwt.Token) (interface{}, error) {
//			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
//				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
//			}
//			return []byte(SecretKey), nil
//		})
//	if err != nil {
//		return false
//	}
//
//	if !token.Valid {
//		return false
//	}
//
//	return true
//}

func GetUserIDFromToken(tokenString string) (uint, error) {
	// создаём экземпляр структуры с утверждениями

	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(SecretKey), nil
		})
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return claims.UserID, nil
}
