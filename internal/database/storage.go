package database

import (
	"errors"
	"strings"

	"gorm.io/gorm"

	"github.com/al-kirpichenko/gophkeeper/internal/models"
)

// ErrConflict - для ошибок вставки
var ErrConflict = errors.New("conflict on inserting new record")

// CreateSecret - создание новой записи
func CreateSecret(db *gorm.DB, secret *models.Secret) error {

	result := db.Create(&secret)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return ErrConflict
	} else if result.Error != nil {
		return result.Error
	}
	return nil
}

// CreateUser - создание пользователя
func CreateUser(db *gorm.DB, user *models.User) error {

	result := db.Create(&user)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return ErrConflict
	} else if result.Error != nil {
		return result.Error
	}
	return nil
}
