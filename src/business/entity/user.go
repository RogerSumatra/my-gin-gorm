package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string  `json:"name" gorm:"type:varchar(255)"`
	Username    string  `json:"username" gorm:"type:varchar(50); UNIQUE; NOT NULL"`
	ShortName   string  `json:"shortName" gorm:"type:varchar(10)"`
	Email       string  `json:"email" gorm:"type:varchar(50); UNIQUE; NOT NULL"`
	Password    string  `json:"password" gorm:"type:varchar(250); NOT NULL"`
	Picture     string  `json:"picture" gorm:"varchar(255)"` //berupa link foto. saat signup foto default kayak no profile
	PhoneNumber string  `json:"phoneNumber" gorm:"type:varchar(50)"`
	Balance     float32 `json:"balance"`
}

type UserBody struct {
	Name        string `json:"name" binding:"required"`
	ShortName   string `json:"shortName" binding:"required"`
	Picture     string `json:"picture" binding:"required"` //berupa link foto.
	PhoneNumber string `json:"phoneNumber" binding:"required"`
}

type UserParam struct {
	UserID int64 `uri:"user_id" gorm:"column:id"`
	PaginationParam
}
