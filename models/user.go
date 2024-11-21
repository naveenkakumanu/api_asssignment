package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	RefreshToken string `json:"refresh_token"`
}
