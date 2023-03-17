package entity

var (
	MalangRegency   = "Malang" //1
	SurabayaRegency = "Surabaya" //2
	BatuRegency     = "Batu" //3
	SemarangRegency = "Semarang" //4
	SoloRegency     = "Solo" //5
	MagelangRegency = "Magelang" //6
	BandungRegency  = "Bandung" //7
	CimahiRegency   = "Cimahi" //8
	BekasiRegency   = "Bekasi" //9
)

type Regency struct {
	ID   uint
	Name string `gorm:"type:varchar(30)"`
}
