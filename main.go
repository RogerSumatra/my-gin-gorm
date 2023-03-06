package main

import (
	"fmt"
	"my-gin-gorm/sdk/config"
	"my-gin-gorm/database"
	"my-gin-gorm/src/handler"
)

func main() {

	

	conf := config.Init()
	if err := conf.Load(".env"); err != nil {
		panic(err)
	}

	sql, err := database.Init(conf)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected successfully")

	handler := handler.Init(conf, sql.GetInstance())
	handler.Run()

}
