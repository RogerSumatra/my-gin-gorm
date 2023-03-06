package entity

import "gorm.io/gorm"

//selfnote = tambah column variabel Post dan PostBody harus sama, format sesuai masing-masing

type Studio struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(255)"`
	Description string `json:"description" gorm:"type:text"`
}

type StudioBody struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type StudioParam struct {
	StudioID int64 `uri:"studio_id" gorm:"column:id"`
	PaginationParam
}
