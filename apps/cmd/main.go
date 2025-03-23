package main

import (
	"go-gin-sample/apps/config"
	"go-gin-sample/apps/controllers"
	"go-gin-sample/apps/db"
	"go-gin-sample/apps/repository"
	"go-gin-sample/apps/services"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()

	db := db.SetupDB()

	itemRepository := repository.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	r := gin.Default()
	r.GET("/items", itemController.FindAll)
	r.GET("items/:id", itemController.FindById)
	r.POST("/items", itemController.Create)
	r.PUT("/items/:id", itemController.Update)
	r.DELETE("/items/:id", itemController.Delete)

	r.Run("localhost:8080")
}
