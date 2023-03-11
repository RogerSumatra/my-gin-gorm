package entity

import "gorm.io/gorm"

type Studio struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(255); NOT NULL"`
	Price float64 `json:"price"`
	Regency  string    `json:"regency" gorm:"type:varchar(255); NOT NULL"`
	City     string    `json:"city" gorm:"type:varchar(255) NOT NULL"`
	//Picture  string    `json:"picture" gorm:"varchar(255)"` //link foto. supabase bingung
	Facility Facility  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`//1 studio 1 facility
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`//1 studio banyak comment
}

type StudioBody struct {
	Name    string `json:"name" binding:"required"`
	Regency string `json:"regency" binding:"required"`
	City    string `json:"city" binding:"required"`
	Price float64 `json:"price" binding:"required"`
	//Picture string `json:"picture" binding:"required"` //link foto. supabase bingung
}

type StudioParam struct {
	StudioID int64 `uri:"studio_id" gorm:"column:id"`
	PaginationParam
}
