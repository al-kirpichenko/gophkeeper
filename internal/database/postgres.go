package database

import (
	"log"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

func InitDB(conf string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(conf), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	return db
}
