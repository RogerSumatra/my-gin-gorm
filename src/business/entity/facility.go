package entity

import "gorm.io/gorm"

type Facility struct {
	gorm.Model
	Picture  string `json:"picture" gorm:"varchar(255)"` //berupa link foto. supabase bingung
	Content  string `json:"content" gorm:"type:text"`
	StudioID uint 

}

type FacilityBody struct {
	StudioID uint `json:"studio_id" binding:"required"`
	Content string `json:"content" binding:"required"`
	//Picture  string `json:"picture" binding:"required"` //berupa link foto. supabase bingung.
}

type FacilityParam struct {
	FacilityID int64 `uri:"facility_id" gorm:"column:id"`
	PaginationParam
}
