package database

import (
	"log"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"

	"github.com/al-kirpichenko/gophkeeper/internal/models"
)

func InitDB(conf string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(conf), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	errMigrate := db.AutoMigrate(&models.User{}, &models.Secret{})

	if errMigrate != nil {
		log.Println("Failed migrate to database!")
		log.Fatal(err)
		return nil
	}

	return db
}
