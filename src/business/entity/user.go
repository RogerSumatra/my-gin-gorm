package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(50);UNIQUE"`
	Password string `json:"password" gorm:"type:varchar(255)"`
}
