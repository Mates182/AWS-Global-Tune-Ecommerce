package router

import (
	"list_shipping_details/internal/repository"
	"list_shipping_details/internal/service"
	"list_shipping_details/internal/controller"
	"github.com/gin-gonic/gin"
	"list_shipping_details/config/cors"
)

func SetupRouter(serviceId string) *gin.Engine {
	router := gin.Default()
	router.Use(cors.GetCORSConfig())
	//[ginshot-routes]
	router.GET("/list", controller.NewListShippingDetailsController(service.NewListShippingDetailsServiceImpl(dbcontext.GetDBClient(), serviceId)).ListShippingDetails)
	//[HttpGET] Ping to list_shipping_details API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
	