package storage

import (
	"errors"

	"gorm.io/gorm"

	"github.com/al-kirpichenko/gophkeeper/internal/database"
	"github.com/al-kirpichenko/gophkeeper/internal/models"
)

// ErrConflict - для ошибок вставки
var (
	ErrConflict     = errors.New("duplicate key value violates unique")
	ErrUserNotFound = errors.New("user not found")
)

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
func (s *Storage) CreateUser(user *models.User) error {

	result := s.db.Create(user)

	if result.Error != nil && errors.Is(result.Error, ErrConflict) {
		return ErrConflict
	} else if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetUser - возвращает модель User по логину
func (s *Storage) GetUser(login string) (*models.User, error) {

	var user *models.User

	result := s.db.First(&user, "login = ?", login)

	if result.Error != nil && errors.Is(result.Error, ErrUserNotFound) {
		return nil, ErrUserNotFound
	}

	return user, nil
}
