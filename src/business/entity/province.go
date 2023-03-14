package entity

var (
	JawaTimurProvince  = "Jawa Timur"
	JawaTengahProvince = "Jawa Tengah"
	JawaBaratProvince  = "Jawa Barat"
)

type Province struct {
	ID   uint
	Name string `gorm:"type:varchar(30)"`
}
