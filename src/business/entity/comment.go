package entity

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	StudioID uint //`json:"studio_id"`
	//Studio   Studio
	UserID  uint    //`json:"user_id"`
	User    User    //belongs to
	Content string  `json:"content" gorm:"type:text"`
	Rating  float32 `json:"rating"`
}

type CommentBody struct {
	StudioID uint    `json:"studio_id" binding:"required"`
	Content  string  `json:"content" binding:"required"`
	Rating   float32 `json:"rating" binding:"required"`
}

type CommentParam struct {
	CommentID int64 `uri:"comment_id" gorm:"column:id"`
	PaginationParam
}
