package router

import (
	"delete_shipping_details/internal/repository"
	"delete_shipping_details/internal/service"
	"delete_shipping_details/internal/controller"
	"github.com/gin-gonic/gin"
	"delete_shipping_details/config/cors"
)

func SetupRouter(serviceId string) *gin.Engine {
	router := gin.Default()
	router.Use(cors.GetCORSConfig())
	//[ginshot-routes]
	router.DELETE("/delete/:id", controller.NewDeleteShippingDetailsController(service.NewDeleteShippingDetailsServiceImpl(dbcontext.GetDBClient(), serviceId)).DeleteShippingDetails)
	//[HttpGET] Ping to delete_shipping_details API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
	