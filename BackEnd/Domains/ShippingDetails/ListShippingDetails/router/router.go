package router

import (
	"list-shipping-details/dbcontext/shippingdetails"
	"list-shipping-details/service"
	"list-shipping-details/controller"
	"github.com/gin-gonic/gin"
	"list-shipping-details/config/cors"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.GetCORSConfig())
	//[ginshot-routes]
	router.GET("/list", controller.NewListShippingDetailsController(service.NewListShippingDetailsServiceImpl(dbcontext.GetDBClient())).ListShippingDetails)
	//[HttpGET] Ping to list-shipping-details API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
	