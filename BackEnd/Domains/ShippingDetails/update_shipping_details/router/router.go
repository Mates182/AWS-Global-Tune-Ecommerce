package router

import (
	"update_shipping_details/internal/repository"
	"update_shipping_details/internal/service"
	"update_shipping_details/internal/controller"
	"github.com/gin-gonic/gin"
	"update_shipping_details/config/cors"
)

func SetupRouter(serviceId string) *gin.Engine {
	router := gin.Default()
	router.Use(cors.GetCORSConfig())
	//[ginshot-routes]
	router.PATCH("/update", controller.NewUpdateShippingDetailsController(service.NewUpdateShippingDetailsServiceImpl(dbcontext.GetDBClient(), serviceId)).UpdateShippingDetails)
	//[HttpGET] Ping to update_shipping_details API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
	