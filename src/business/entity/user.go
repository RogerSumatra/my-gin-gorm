package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(50); UNIQUE; NOT NULL"`
	ShortName string `json:"shortName" gorm:"type:varchar(10)"`
	Email    string `json:"email" gorm:"type:varchar(50); UNIQUE; NOT NULL"`
	Password string `json:"password" gorm:"type:varchar(250); NOT NULL"`
	//Picture  string `json:"picture" gorm:"varchar(255)"` //berupa link foto. supabase bingung. Atau default kayak no profile
	Balance float32 `json:"balance"`
}

type UserParam struct {
	UserID int64 `uri:"user_id" gorm:"column:id"`
	PaginationParam
}
