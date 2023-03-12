package entity

type City struct {
	ID   uint   
	Name string `gorm:"type:varchar(30)"`
}
