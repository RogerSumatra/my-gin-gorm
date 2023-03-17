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
			Duration: entity.Duration{
				CheckIn:  entity.CheckIn9,
				CheckOut: entity.CheckOut10,
			},
		},
		{
			Hour: entity.Set10to11,
			Duration: entity.Duration{
				CheckIn:  entity.CheckIn10,
				CheckOut: entity.CheckOut11,
			},
		},
		{
			Hour: entity.Set11to12,
			Duration: entity.Duration{
				CheckIn:  entity.CheckIn11,
				CheckOut: entity.CheckOut12,
			},
		},
		{
			Hour: entity.Set12to13,
			Duration: entity.Duration{
				CheckIn:  entity.CheckIn12,
				CheckOut: entity.CheckOut13,
			},
		},
		{
			Hour: entity.Set13to14,
			Duration: entity.Duration{
				CheckIn:  entity.CheckIn13,
				CheckOut: entity.CheckOut14,
			},
		},
		{
			Hour: entity.Set15to16,
			Duration: entity.Duration{
				CheckIn:  entity.CheckIn15,
				CheckOut: entity.CheckOut16,
			},
		},
		{
			Hour: entity.Set16to17,
			Duration: entity.Duration{
				CheckIn:  entity.CheckIn16,
				CheckOut: entity.CheckOut17,
			},
		},
		{
			Hour: entity.Set17to18,
			Duration: entity.Duration{
				CheckIn:  entity.CheckIn17,
				CheckOut: entity.CheckOut18,
			},
		},
		{
			Hour: entity.Set18to19,
			Duration: entity.Duration{
				CheckIn:  entity.CheckIn18,
				CheckOut: entity.CheckOut19,
			},
		},
	}

	if err := sql.Create(&hour).Error; err != nil {
		return err
	}
	return nil
}
