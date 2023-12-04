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

type LP struct {
	gorm.Model
	UserID   string
	Login    string
	Password string
	Comment  string
}

type Text struct {
	gorm.Model
	UserID  string
	Text    string
	Comment string
}

type Binary struct {
	gorm.Model
	UserID  string
	Binary  []byte
	Comment string
}

type Card struct {
	gorm.Model
	UserID  string
	Number  string
	Date    string
	CVV     int
	Comment string
}
