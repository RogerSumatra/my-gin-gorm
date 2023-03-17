package entity

import "gorm.io/gorm"


type History struct {
	//one for every user
	//history has many cart
	gorm.Model
	UserID uint
	User User
	Cart []Cart
	Status string

}



type Review struct {

	gorm.Model
	HistoryID int
	History History
	CartTotal   float64
	BalanceUser float64
}
