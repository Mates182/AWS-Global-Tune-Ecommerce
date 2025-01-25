package router

import (
	"create-order-service/controllers"

	"create-order-service/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(createOrderController *controllers.CreateOrderController) *gin.Engine {
	router := gin.Default()

	router.Use(config.GetCORSConfig())
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("create", createOrderController.CreateOrder)
	return router
}
