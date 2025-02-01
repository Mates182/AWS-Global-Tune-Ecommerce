package router

import (
	"update-shipping-details/dbcontext/shippingdetails"
	"update-shipping-details/service"
	"update-shipping-details/controller"
	"update-shipping-details/config/cors"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.GetCORSConfig())
	//[ginshot-routes]
	router.PATCH("/update", controller.NewUpdateShippingDetailsController(service.NewUpdateShippingDetailsServiceImpl(dbcontext.GetDBClient())).UpdateShippingDetails)
	//[HttpGET] Ping to update-shipping-details API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
