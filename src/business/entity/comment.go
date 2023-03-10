package entity

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	StudioID uint    //`json:"studio_id"`
	Studio Studio 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID   uint    //`json:"user_id"`
	User User //belongs to
	Content  string  `json:"content" gorm:"type:text"`
	Rating   float32 `json:"rating"`
}

type CommentBody struct {
	UserID   uint    `json:"user_id" binding:"required"`
	StudioID uint    `json:"studio_id" binding:"required"`
	Content  string  `json:"content" binding:"required"`
	Rating   float32 `json:"rating" binding:"required"`
}

type CommentParam struct {
	CommentID int64 `uri:"comment_id" gorm:"column:id"`
	PaginationParam
}
