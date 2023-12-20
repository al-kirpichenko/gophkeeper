// Package storage Пакет  содержит функции для сохранения и загрузки токена из файла.
package storage

import (
	"fmt"
	"os"
)

const FilePath = "temp/auth_token.txt"

type TokenService interface {
	SaveToken(token string) error
	LoadToken(token string) error
	GetToken(token string) (string, error)
}

type TokenStorage struct {
	filePath string
}

func New() *TokenStorage {
	return &TokenStorage{
		filePath: FilePath,
	}
}

// SaveToken сохраняет токен в файл.
func (t *TokenStorage) SaveToken(token string) error {

	file, err := os.OpenFile(t.filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	err = os.WriteFile(t.filePath, []byte(token), 0644)
	if err != nil {
		return fmt.Errorf("ошибка при сохранении токена в файл: %v", err)
	}
	return nil
}

// LoadToken загружает токен из файла.
func (t *TokenStorage) LoadToken() (string, error) {
	content, err := os.ReadFile(t.filePath)
	if err != nil {
		return "", fmt.Errorf("ошибка при загрузке токена из файла: %v", err)
	}
	return string(content), nil
}

// GetToken возвращает текущий токен.
func (t *TokenStorage) GetToken() string {
	tokenContent, _ := t.LoadToken()
	return tokenContent
}
