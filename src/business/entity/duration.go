package entity


var (
	CheckIn9  = "09:00"
	CheckIn10 = "10:00"
	CheckIn11 = "11:00"
	CheckIn12 = "12:00"
	CheckIn13 = "13:00"
	CheckIn14 = "14:00"
	CheckIn15 = "15:00"
	CheckIn16 = "16:00"
	CheckIn17 = "17:00"
	CheckIn18 = "18:00"
)

var (
	CheckOut10 = "10:00"
	CheckOut11 = "11:00"
	CheckOut12 = "12:00"
	CheckOut13 = "13:00"
	CheckOut14 = "14:00"
	CheckOut15 = "15:00"
	CheckOut16 = "16:00"
	CheckOut17 = "17:00"
	CheckOut18 = "18:00"
	CheckOut19 = "19:00"
)

type Duration struct {
	ID       uint
	CheckIn  string `gorm:"varchar(20)"`
	CheckOut string `gorm:"varchar(20)"`
}
