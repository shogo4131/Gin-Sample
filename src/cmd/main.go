package main

import (
	"go-gin-sample/apps/config"
	"go-gin-sample/apps/controllers"
	"go-gin-sample/apps/model"
	"go-gin-sample/apps/repository"
	"go-gin-sample/apps/services"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()

	items := []model.Item{
		{Name: "テスト商品", Price: 1000, Description: "テスト商品の説明", SoldOut: false},
		{Name: "テスト商品2", Price: 2000, Description: "テスト商品2の説明", SoldOut: true},
		{Name: "テスト商品3", Price: 3000, Description: "テスト商品3の説明", SoldOut: false},
	}

	itemRepository := repository.NewItemMemoryRepository(items)
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
