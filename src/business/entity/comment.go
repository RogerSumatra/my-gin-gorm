package entity

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	StudioID uint `json:"studio_id"`
	Content string `json:"content"`
	UserID uint `json:"user_id"`
}