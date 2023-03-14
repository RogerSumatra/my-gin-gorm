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