package entity

import "gorm.io/gorm"

//selfnote = tambah column variabel Post dan PostBody harus sama, format sesuai masing-masing

type Studio struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(255)"`
	Description string `json:"description" gorm:"type:text"`
	//Rating  float32 `json:"rating" gorm:"type:decimal(10,2)"`
	//Regency string `json:"regency" gorm:"type:varchar(255)"`
	//City string `json:"alamat" gorm:"type:varchar(255)`
	//Picture string `json:"picture" gorm:"varchar(255)"`
	//Facility Facility `json:"facility" gorm:"type:text"`   //masih dipertanyakan
	Comments []Comment `json:"comments"`
}

type StudioBody struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type StudioParam struct {
	StudioID int64 `uri:"studio_id" gorm:"column:id"`
	PaginationParam
}
