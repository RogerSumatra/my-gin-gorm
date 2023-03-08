package entity

type Facility struct {
	Picture string //foto
	Content string `json:"content" gorm:"type:text"`
}