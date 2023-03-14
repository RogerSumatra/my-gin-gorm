package entity

var (
	MalangRegency   = "Malang"
	SurabayaRegency = "Surabaya"
	BatuRegency     = "Batu"
	SemarangRegency = "Semarang"
	SoloRegency     = "Solo"
	MagelangRegency = "Magelang"
	BandungRegency  = "Bandung"
	CimahiRegency   = "Cimahi"
	BekasiRegency   = "Bekasi"
)

type Regency struct {
	ID   uint
	Name string `gorm:"type:varchar(30)"`
}
