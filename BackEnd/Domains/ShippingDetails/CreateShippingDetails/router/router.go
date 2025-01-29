package router

import (
	"create-shipping-details/dbcontext/shippingdetails"
	"create-shipping-details/service"
	"create-shipping-details/controller"
	"github.com/gin-gonic/gin"
	"create-shipping-details/config/cors"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.GetCORSConfig())
	//[ginshot-routes]
	router.POST("/create", controller.NewCreateShippingDetailsController(service.NewCreateShippingDetailsServiceImpl(dbcontext.GetDBClient())).CreateShippingDetails)
	//[HttpGET] Ping to create-shipping-details API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
	