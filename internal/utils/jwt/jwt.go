package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/al-kirpichenko/gophkeeper/internal/models"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID uint
}

const SecretKey = "bvaEFBtr5e"

// NewToken creates new JWT token for given user and app.
func NewToken(user *models.User, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Добавляем в токен всю необходимую информацию
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["login"] = user.Login
	claims["exp"] = time.Now().Add(duration).Unix()

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
//
//func GetUserIDFromToken(tokenString string) (uint, error) {
//	// создаём экземпляр структуры с утверждениями
//
//	claims := &Claims{}
//	_, err := jwt.ParseWithClaims(tokenString, claims,
//		func(t *jwt.Token) (interface{}, error) {
//			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
//				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
//			}
//			return []byte(SecretKey), nil
//		})
//	if err != nil {
//		log.Println(err)
//		return 0, err
//	}
//	return claims.UserID, nil
//}
