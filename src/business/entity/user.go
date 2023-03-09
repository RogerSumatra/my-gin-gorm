package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(50); UNIQUE; NOT NULL"`
	Email    string `json:"email" gorm:"type:varchar(50); UNIQUE; NOT NULL"`
	Password string `json:"password" gorm:"type:varchar(50); UNIQUE NOT NULL"`
	Picture  string `json:"picture" gorm:"varchar(255)"` //berupa link foto. supabase bingung
	//Saldo
}
