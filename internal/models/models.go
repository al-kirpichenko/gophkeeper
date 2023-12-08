package models

import "gorm.io/gorm"

// Auth - содержит данные для регистрации/ аутентификации
type Auth struct {
	Login    string `json:"login"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

// User - модель пользователя
type User struct {
	gorm.Model
	Login    string `gorm:"size:255;not null;unique"`
	Password []byte `gorm:"size:255;not null"`
}

// Secret - модель хранимого секрета
type Secret struct {
	gorm.Model
	UserID  uint   `gorm:"not null;index:,unique,composite:uid"`
	Title   string `gorm:"size:255;not null;index:,unique,composite:uid" json:"title"`
	Content []byte `gorm:"not null" json:"content"`
	Comment string `gorm:"size:255;not null"`
}
