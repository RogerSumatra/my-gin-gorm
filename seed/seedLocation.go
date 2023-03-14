package seed

import (
	"my-gin-gorm/src/business/entity"

	"gorm.io/gorm"
)

func SeedProvince(sql *gorm.DB) error {
	var province []entity.Province

	if err := sql.First(&province).Error; err != gorm.ErrRecordNotFound {
		return err
	}
	province = []entity.Province{
		{
			Name: entity.JawaTimurProvince,
		},
		{
			Name: entity.JawaTengahProvince,
		},
		{
			Name: entity.JawaBaratProvince,
		},
	}

	if err := sql.Create(&province).Error; err != nil {
		return err
	}
	return nil
}

func SeedRegency(sql *gorm.DB) error {
	var regency []entity.Regency

	if err := sql.First(&regency).Error; err != gorm.ErrRecordNotFound {
		return err
	}
	regency = []entity.Regency{
		{
			Name: entity.MalangRegency,
		},
		{
			Name: entity.SurabayaRegency,
		},
		{
			Name: entity.BatuRegency,
		},
		{
			Name: entity.SemarangRegency,
		},
		{
			Name: entity.SoloRegency,
		},
		{
			Name: entity.MagelangRegency,
		},
		{
			Name: entity.BandungRegency,
		},
		{
			Name: entity.CimahiRegency,
		},
		{
			Name: entity.BekasiRegency,
		},		
	}

	if err := sql.Create(&regency).Error; err != nil {
		return err
	}
	return nil
}