package entity

import "gorm.io/gorm"

type Studio struct {
	gorm.Model
	Name     string     `json:"name" gorm:"type:varchar(255); NOT NULL"`
	Price    float64    `json:"price"`
	Regency  []Regency  `gorm:"many2many:studio_regency;"`
	Province []Province `gorm:"many2many:studio_province;"`
	//Picture  string    `json:"picture" gorm:"varchar(255)"` //link foto
	Facility Facility  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` //1 studio 1 facility
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` //1 studio banyak comment
}

type StudioBody struct {
	Name       string  `json:"name" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
	ProvinceID []uint  `json:"provinceID"`
	RegencyID  []uint  `json:"regencyID"`
	//Picture string `json:"picture" binding:"required"` //link foto
}

type StudioParam struct {
	StudioID int64 `uri:"studio_id" gorm:"column:id"`
	PaginationParam
}
