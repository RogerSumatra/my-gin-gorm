package entity

import "gorm.io/gorm"

type Studio struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(255); NOT NULL"`
	//Harga
	Regency  string    `json:"regency" gorm:"type:varchar(255); NOT NULL"`
	City     string    `json:"city" gorm:"type:varchar(255) NOT NULL"`
	Picture  string    `json:"picture" gorm:"varchar(255)"` //link foto. supabase bingung
	Facility Facility  //1 studio 1 facility
	Comments []Comment //1 studio banyak comment
}

type StudioBody struct { //rencana untuk di list
	Name    string `json:"name" binding:"required"`
	Regency string `json:"regency" binding:"required"`
	City    string `json:"city" binding:"required"`
	Picture string `json:"picture" binding:"required"` //link foto. supabase bingung
	//Harga
}

type StudioParam struct {
	StudioID int64 `uri:"studio_id" gorm:"column:id"`
	PaginationParam
}
