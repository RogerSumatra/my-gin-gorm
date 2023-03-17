package entity

var (
	Set9to10  = "09:00-10:00"
	Set10to11 = "10:00-11:00"
	Set11to12 = "11:00-12:00"
	Set12to13 = "12:00-13:00"
	Set13to14 = "13:00-14:00"
	Set14to15 = "14:00-15:00"
	Set15to16 = "15:00-16:00"
	Set16to17 = "16:00-17:00"
	Set17to18 = "17:00-18:00"
	Set18to19 = "18:00-19:00"
)
type Hour struct {
	ID       uint
	Hour     string `gorm:"varchar(10)"`
	DurationID uint
	Duration Duration
}

type HourParam struct {
	HourID int64 `uri:"hour_id" gorm:"column:id"`
	PaginationParam
}
