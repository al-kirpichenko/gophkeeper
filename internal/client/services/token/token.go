// Package token содержит функции для работы с токеном.
package token

//import (
//	"fmt"
//
//	"github.com/al-kirpichenko/gophkeeper/internal/client/services/token"
//)
//
//const tokenFilePath = "auth_token.txt"
//

//// IsAvailable проверяет наличие токена.
//func IsAvailable() bool {
//	content, err := token.LoadToken()
//	return err == nil && content != ""
//}

//// CheckValidity проверяет наличие и годность токена.
//func CheckValidity() error {
//	// Проверка наличия токена
//	if !IsAvailable() {
//		fmt.Println("Токен отсутствует. Выполните аутентификацию командой 'login'.")
//		return fmt.Errorf("токен отсутствует")
//	}
//
//	// Проверка годности токена
//	if !IsValid() {
//		fmt.Println("Токен устарел. Выполните аутентификацию командой 'login'.")
//		return fmt.Errorf("токен устарел")
//	}
//
//	return nil
//}

//// IsValid проверяет годность токена.
//func IsValid() bool {
//	// Здесь можно добавить дополнительную логику проверки годности токена,
//	// например, проверку срока действия.
//	return true
//}
//
//// FilePath возвращает путь к файлу токена.
//func FilePath() string {
//	return tokenFilePath
//}
