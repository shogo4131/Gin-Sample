package main

import (
	"go-gin-sample/apps/config"
	"go-gin-sample/apps/db"
	"go-gin-sample/apps/model"
)

func main() {
	config.InitConfig()

	db := db.SetupDB()

	if err := db.AutoMigrate(&model.Item{}); err != nil {
		panic("failed to migrate")
	}
}
