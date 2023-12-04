package models

import "gorm.io/gorm"

type Auth struct {
	Login    string `json:"login"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type User struct {
	gorm.Model
	Login    string `gorm:"size:255;not null;unique"`
	Password string `gorm:"size:255;not null"`
}

type Secret struct {
	gorm.Model
	UserID string `gorm:"size:255;not null"`
	Text   string `gorm:"not null"`
}
