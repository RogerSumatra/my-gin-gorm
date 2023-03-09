package entity

import "gorm.io/gorm"

type Facility struct {
	gorm.Model
	Picture  string `json:"picture" gorm:"varchar(255)"` //berupa link foto. supabase bingung
	Content  string `json:"content" gorm:"type:text"`
	StudioID uint 
}
