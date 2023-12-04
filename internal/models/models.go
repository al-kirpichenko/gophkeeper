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

type Psw struct {
	gorm.Model
	UserID   string `gorm:"size:255;not null"`
	Login    string `gorm:"size:255;not null"`
	Password string `gorm:"size:255;not null"`
	Comment  string `gorm:"size:255;not null"`
}

type Text struct {
	gorm.Model
	UserID  string `gorm:"size:255;not null"`
	Text    string `gorm:"size:3000;not null"`
	Comment string `gorm:"size:255;not null"`
}

type Binary struct {
	gorm.Model
	UserID  string `gorm:"size:255;not null"`
	Binary  []byte `gorm:"not null"`
	Comment string `gorm:"size:255;not null"`
}

type Card struct {
	gorm.Model
	UserID  string `gorm:"size:255;not null"`
	Number  string `gorm:"size:22;not null"`
	Date    string `gorm:"size:8;not null"`
	CVV     int    `gorm:"size:3;not null"`
	Comment string
}
