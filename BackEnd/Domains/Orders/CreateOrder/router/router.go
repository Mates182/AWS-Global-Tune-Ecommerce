package router

import (
	"create-order-service/controllers"

	"create-order-service/config"

	"github.com/gin-gonic/gin"
)

func SetupRouter(createOrderController *controllers.CreateOrderController) *gin.Engine {
	router := gin.Default()

	router.Use(config.GetCORSConfig())

	router.POST("create", createOrderController.CreateOrder)
	return router
}
