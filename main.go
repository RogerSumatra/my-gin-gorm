package main

import (
	"fmt"
	"my-gin-gorm/sdk/config"
	"my-gin-gorm/sdk/sql"
	"my-gin-gorm/src/business/entity"
	"my-gin-gorm/src/handler"
)

func main() {
	conf := config.Init()
	if err := conf.Load(".env"); err != nil {
		panic(err)
	}
	sqlConfig := sql.Config{
		Username: conf.Get("DB_USERNAME"),
		Password: conf.Get("DB_PASSWORD"),
		Host:     conf.Get("DB_HOST"),
		Port:     conf.Get("DB_PORT"),
		Database: conf.Get("DB_DATABASE"),
	}
	sql, err := sql.Init(sqlConfig)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected successfully")

	db := sql.GetInstance()
	db.AutoMigrate(entity.Post{}) //tabel data studio (kayaknya)
	db.AutoMigrate(entity.User{}) //tabel akun

	handler := handler.Init(conf, db)
	handler.Run()

}
