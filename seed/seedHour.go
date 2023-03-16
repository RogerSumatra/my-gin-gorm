package seed

import (
	"my-gin-gorm/src/business/entity"

	"gorm.io/gorm"
)

func SeedHour(sql *gorm.DB) error {
	var hour []entity.Hour

	if err := sql.First(&hour).Error; err != gorm.ErrRecordNotFound {
		return err
	}
	hour = []entity.Hour{
		{
			Hour: entity.Set9to10,
		},
		{
			Hour: entity.Set10to11,
		},
		{
			Hour: entity.Set11to12,
		},
		{
			Hour: entity.Set12to13,
		},
		{
			Hour: entity.Set13to14,
		},
		{
			Hour: entity.Set14to15,
		},
		{
			Hour: entity.Set15to16,
		},
		{
			Hour: entity.Set16to17,
		},
		{
			Hour: entity.Set17to18,
		},
		{
			Hour: entity.Set18to19,
		},
	}

	if err := sql.Create(&hour).Error; err != nil {
		return err
	}
	return nil
}
