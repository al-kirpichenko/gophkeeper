package database

import (
	"gorm.io/gorm"

	"github.com/al-kirpichenko/gophkeeper/internal/models"
)

func CreateSecret(db *gorm.DB, secret *models.Secret) error {

	result := db.Create(&secret)

	if result.Error != nil {
		return result.Error
	}
	return nil

}
