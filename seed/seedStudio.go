package seed

import (
	"my-gin-gorm/src/business/entity"

	"gorm.io/gorm"
)

func SeedStudio(sql *gorm.DB) error {
	var studio []entity.Studio
	//var regency []entity.Regency
	//var province []entity.Province

	if err := sql.First(&studio).Error; err != gorm.ErrRecordNotFound {
		return err
	}
	// if err := sql.First(&regency).Error; err != gorm.ErrRecordNotFound {
	// 	return err
	// }

	studio = []entity.Studio{
		{
			Name:  "Pixel Playground",
			Price: 150000,
			Regency: []entity.Regency{
				{
					ID:   1,
					Name: entity.MalangRegency,
				},
			},
			Province: []entity.Province{
				{
					ID:   1,
					Name: entity.JawaTimurProvince,
				},
			},
			Picture:  "https://prtlclzsqdfetjatgvbw.supabase.co/storage/v1/object/public/api-service/studio-foto_169.jpeg",
			Discount: 25,
		},
		{
			Name:  "Creative Canvas",
			Price: 10000,
			Regency: []entity.Regency{
				{
					ID:   7,
					Name: entity.BandungRegency,
				},
			},
			Province: []entity.Province{
				{
					ID:   1,
					Name: entity.JawaTimurProvince,
				},
			},
			Picture:  "https://prtlclzsqdfetjatgvbw.supabase.co/storage/v1/object/public/api-service/studio-foto_169.jpeg",
			Discount: 10,
		},
		{
			Name:  "Stellar Studios",
			Price: 17500,
			Regency: []entity.Regency{
				{
					ID:   5,
					Name: entity.SoloRegency,
				},
			},
			Province: []entity.Province{
				{
					ID:   2,
					Name: entity.JawaTengahProvince,
				},
			},
			Picture:  "https://prtlclzsqdfetjatgvbw.supabase.co/storage/v1/object/public/api-service/studio-foto_169.jpeg",
			Discount: 25,
		},
		{
			Name:  "Bluebird Productions",
			Price: 20000,
			Regency: []entity.Regency{
				{
					ID:   9,
					Name: entity.BekasiRegency,
				},
			},
			Province: []entity.Province{
				{
					ID:   3,
					Name: entity.JawaBaratProvince,
				},
			},
			Picture:  "https://prtlclzsqdfetjatgvbw.supabase.co/storage/v1/object/public/api-service/studio-foto_169.jpeg",
			Discount: 10,
		},
		{
			Name:  "Bright Idea Studios",
			Price: 52000,
			Regency: []entity.Regency{
				{
					ID:   9,
					Name: entity.BekasiRegency,
				},
			},
			Province: []entity.Province{
				{
					ID:   3,
					Name: entity.JawaBaratProvince,
				},
			},
			Picture:  "https://prtlclzsqdfetjatgvbw.supabase.co/storage/v1/object/public/api-service/studio-foto_169.jpeg",
			Discount: 10,
		},
		{
			Name:  "Emerald Eye Studios",
			Price: 20000,
			Regency: []entity.Regency{
				{
					ID:   4,
					Name: entity.SemarangRegency,
				},
			},
			Province: []entity.Province{
				{
					ID:   2,
					Name: entity.JawaTengahProvince,
				},
			},
			Picture:  "https://prtlclzsqdfetjatgvbw.supabase.co/storage/v1/object/public/api-service/studio-foto_169.jpeg",
			Discount: 30,
		},
		{
			Name:  "Golden Hour Productions",
			Price: 43000,
			Regency: []entity.Regency{
				{
					ID:   8,
					Name: entity.CimahiRegency,
				},
			},
			Province: []entity.Province{
				{
					ID:   3,
					Name: entity.JawaBaratProvince,
				},
			},
			Picture:  "https://prtlclzsqdfetjatgvbw.supabase.co/storage/v1/object/public/api-service/studio-foto_169.jpeg",
			Discount: 15,
		},
		{
			Name:  "Harmoni Studio",
			Price: 24000,
			Regency: []entity.Regency{
				{
					ID:   2,
					Name: entity.SurabayaRegency,
				},
			},
			Province: []entity.Province{
				{
					ID:   1,
					Name: entity.JawaTimurProvince,
				},
			},
			Picture:  "https://prtlclzsqdfetjatgvbw.supabase.co/storage/v1/object/public/api-service/studio-foto_169.jpeg",
			Discount: 10,
		},
		{
			Name:  "Rhythm Room",
			Price: 80000,
			Regency: []entity.Regency{
				{
					ID:   9,
					Name: entity.BekasiRegency,
				},
			},
			Province: []entity.Province{
				{
					ID:   3,
					Name: entity.JawaBaratProvince,
				},
			},
			Picture:  "https://prtlclzsqdfetjatgvbw.supabase.co/storage/v1/object/public/api-service/studio-foto_169.jpeg",
			Discount: 35,
		},
		{
			Name:  "Melody Maker",
			Price: 24000,
			Regency: []entity.Regency{
				{
					ID:   3,
					Name: entity.BatuRegency,
				},
			},
			Province: []entity.Province{
				{
					ID:   1,
					Name: entity.JawaTimurProvince,
				},
			},
			Picture:  "https://prtlclzsqdfetjatgvbw.supabase.co/storage/v1/object/public/api-service/studio-foto_169.jpeg",
			Discount: 22,
		},
		{
			Name:  "Sound Lab",
			Price: 43000,
			Regency: []entity.Regency{
				{
					ID:   6,
					Name: entity.MagelangRegency,
				},
			},
			Province: []entity.Province{
				{
					ID:   2,
					Name: entity.JawaTengahProvince,
				},
			},
			Picture:  "https://prtlclzsqdfetjatgvbw.supabase.co/storage/v1/object/public/api-service/studio-foto_169.jpeg",
			Discount: 13,
		},
		{
			Name:  "Beat Box Studio",
			Price: 22000,
			Regency: []entity.Regency{
				{
					ID:   7,
					Name: entity.BandungRegency,
				},
			},
			Province: []entity.Province{
				{
					ID:   1,
					Name: entity.JawaTimurProvince,
				},
			},
			Picture:  "https://prtlclzsqdfetjatgvbw.supabase.co/storage/v1/object/public/api-service/studio-foto_169.jpeg",
			Discount: 21,
		},
		{
			Name:  "Note-worthy Studio",
			Price: 32000,
			Regency: []entity.Regency{
				{
					ID:   4,
					Name: entity.SemarangRegency,
				},
			},
			Province: []entity.Province{
				{
					ID:   2,
					Name: entity.JawaTengahProvince,
				},
			},
			Picture:  "https://prtlclzsqdfetjatgvbw.supabase.co/storage/v1/object/public/api-service/studio-foto_169.jpeg",
			Discount: 10,
		},
		{
			Name:  "Groove House",
			Price: 15500,
			Regency: []entity.Regency{
				{
					ID:   9,
					Name: entity.BekasiRegency,
				},
			},
			Province: []entity.Province{
				{
					ID:   3,
					Name: entity.JawaBaratProvince,
				},
			},
			Picture:  "https://prtlclzsqdfetjatgvbw.supabase.co/storage/v1/object/public/api-service/studio-foto_169.jpeg",
			Discount: 10,
		},
		{
			Name:  "Music Mill Studio",
			Price: 15000,
			Regency: []entity.Regency{
				{
					ID:   9,
					Name: entity.BekasiRegency,
				},
			},
			Province: []entity.Province{
				{
					ID:   3,
					Name: entity.JawaBaratProvince,
				},
			},
			Picture:  "https://prtlclzsqdfetjatgvbw.supabase.co/storage/v1/object/public/api-service/studio-foto_169.jpeg",
			Discount: 15,
		},
	}

	if err := sql.Create(&studio).Error; err != nil {
		return err
	}
	return nil
}
