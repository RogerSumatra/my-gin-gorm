package entity

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	StudioID uint    
	Content  string  `json:"content" gorm:"type:text"`
	UserID   uint    
	Rating   float32 `json:"rating"`
}
