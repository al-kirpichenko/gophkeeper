package storage

import (
	"errors"
	"strings"

	"gorm.io/gorm"

	"github.com/al-kirpichenko/gophkeeper/internal/database"
	"github.com/al-kirpichenko/gophkeeper/internal/models"
)

// ErrConflict - для ошибок вставки
var ErrConflict = errors.New("conflict on inserting new record")

type Storage struct {
	db *gorm.DB
}

// NewStorage - конструктор хранилища
func NewStorage(conf string) *Storage {
	return &Storage{
		db: database.InitDB(conf),
	}
}

// CreateUser - создает нового пользователья в бд
func (s *Storage) CreateUser(auth *models.Auth) error {

	result := s.db.Create(auth)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return ErrConflict
	} else if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetUser - возвращает модель User по логину
func (s *Storage) GetUser(login string) (*models.User, error) {
	return &models.User{}, nil
}
