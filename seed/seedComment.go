package seed

import (
	"my-gin-gorm/src/business/entity"

	"gorm.io/gorm"
)

func SeedComment(sql *gorm.DB) error {
	var comment []entity.Comment
	//var regency []entity.Regency
	//var province []entity.Province

	if err := sql.First(&comment).Error; err != gorm.ErrRecordNotFound {
		return err
	}
	// if err := sql.First(&regency).Error; err != gorm.ErrRecordNotFound {
	// 	return err
	// }

	comment = []entity.Comment{
		{

		},
	}

	if err := sql.Create(&comment).Error; err != nil {
		return err
	}
	return nil
}
