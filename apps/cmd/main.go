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

	authRepository := repository.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	r := gin.Default()
	itemRouter := r.Group("/items")
	authRouter := r.Group("/auth")

	itemRouter.GET("", itemController.FindAll)
	itemRouter.GET("/:id", itemController.FindById)
	itemRouter.POST("", itemController.Create)
	itemRouter.PUT("/:id", itemController.Update)
	itemRouter.DELETE("/:id", itemController.Delete)

	authRouter.POST("/signup", authController.Signup)

	r.Run("localhost:8080")
}
