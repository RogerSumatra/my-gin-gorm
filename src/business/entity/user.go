package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"type:varchar(50);UNIQUE"`
	Password string `json:"password" gorm:"type:varchar(255)"`
}
