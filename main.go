package main

import (
	"my-gin-gorm/sdk/config"
)

func main() {
	conf := config.Init()
	if err := conf.Load(".env"); err != nil {
		panic(err)
	}
}
