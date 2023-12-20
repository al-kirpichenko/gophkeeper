package storage

import (
	"errors"

	"gorm.io/gorm"

	"github.com/al-kirpichenko/gophkeeper/internal/database"
	"github.com/al-kirpichenko/gophkeeper/internal/models"
)

// ErrConflict - для ошибок вставки
var (
	ErrConflict = errors.New("duplicate key value violates unique")
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

// CreateUser  - создает нового пользователья в бд
func (s *Storage) CreateUser(user *models.User) error {

	result := s.db.Create(user)

	if result.Error != nil && errors.Is(result.Error, ErrConflict) {
		return ErrConflict
	} else if result.Error != nil {
		return result.Error
	}
	return nil
}

// ReadUser - возвращает модель User по логину
func (s *Storage) ReadUser(login string) (*models.User, error) {

	var user *models.User

	result := s.db.First(&user, "login = ?", login)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

// CreateSecret - создает новый секрет пользователя
func (s *Storage) CreateSecret(secret *models.Secret) error {

	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(secret).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil && errors.Is(err, ErrConflict) {
		return ErrConflict
	} else if err != nil {
		return err
	}
	return nil
}

func (s *Storage) ReadSecret(title string, uid uint) (*models.Secret, error) {

	var secret *models.Secret

	result := s.db.Where("title = ? AND user_id = ?", title, uid).Find(&secret)

	if result.Error != nil {
		return nil, result.Error
	}

	return secret, nil
}
