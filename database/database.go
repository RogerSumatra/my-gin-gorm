package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"my-gin-gorm/sdk/config"
	"my-gin-gorm/seed"
	"my-gin-gorm/src/business/entity"
)

type Interface interface {
	GetInstance() *gorm.DB
}

type Config struct {
	Host     string
	Port     string
	Password string
	Username string
	Database string
}

type sql struct {
	Db *gorm.DB
}

func Init(conf config.Interface) (Interface, error) {

	sql := sql{}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Get("DB_USERNAME"),
		conf.Get("DB_PASSWORD"),
		conf.Get("DB_HOST"),
		conf.Get("DB_PORT"),
		conf.Get("DB_DATABASE"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(entity.User{}, entity.Studio{}, &entity.Comment{}, &entity.Facility{}, &entity.Province{}, &entity.Regency{})
	sql.Db = db

	if err := seed.SeedProvince(sql.Db); err != nil {
		panic("Set Province Seed Failed")
	}
	if err := seed.SeedRegency(sql.Db); err != nil {
		panic("Set Regency Seed Failed")
	}

	return &sql, nil
}

func (s *sql) GetInstance() *gorm.DB {
	return s.Db
}
