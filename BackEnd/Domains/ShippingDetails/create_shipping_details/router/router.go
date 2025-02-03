package router

import (
	"create_shipping_details/internal/repository"
	"create_shipping_details/internal/service"
	"create_shipping_details/internal/controller"
	"github.com/gin-gonic/gin"
	"create_shipping_details/config/cors"
)

func SetupRouter(serviceId string) *gin.Engine {
	router := gin.Default()
	router.Use(cors.GetCORSConfig())
	//[ginshot-routes]
	router.POST("/create", controller.NewCreateShippingDetailsController(service.NewCreateShippingDetailsServiceImpl(dbcontext.GetDBClient(), serviceId)).CreateShippingDetails)
	//[HttpGET] Ping to create_shipping_details API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
	