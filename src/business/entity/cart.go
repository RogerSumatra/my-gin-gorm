package entity

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	NameContact string `json:"nameContact"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"iniEmail"`
	UserID      uint
	User        User
	StudioID    uint
	Studio      Studio
	HourID      uint
	Hour        Hour
}

type CartRespond struct { //nampilin
	StudioID       uint
	StudioTax      float64
	StudioCheckIn  string
	StudioCheckOut string
	StudioPrice    float64
	Total          float64
}

type CartRequest struct {
	NameContact string `json:"nameContact" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Email       string `json:"iniEmail" binding:"required"`
}

type CartParam struct {
	CartID int64 `uri:"cart_id" gorm:"column:id"`
	PaginationParam
}
