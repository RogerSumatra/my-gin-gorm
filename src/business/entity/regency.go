package entity

type Regency struct {
	ID   uint   
	Name string `gorm:"type:varchar(30)"`
}
