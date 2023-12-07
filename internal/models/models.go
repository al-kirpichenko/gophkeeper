package models

import "gorm.io/gorm"

type Auth struct {
	Login    string `json:"login"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserSecret struct {
	Name       string `json:"name"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Text       string `json:"text"`
	Binary     []byte `json:"binary"`
	CardNumber string `json:"card_number"`
	Expiry     string `json:"expiry"`
	Bank       string `json:"bank"`
	CVV        string `json:"cvv"`
}

type User struct {
	gorm.Model
	Login    string `gorm:"size:255;not null;unique"`
	Password string `gorm:"size:255;not null"`
}

type Secret struct {
	gorm.Model
	UserID string `gorm:"size:255;not null"`
	Name   string `gorm:"size:255;not null;unique_index:user_secret"`
	Text   string `gorm:"not null"`
}
