package router

import (
	"get_shipping_details_by_id/internal/repository"
	"get_shipping_details_by_id/internal/service"
	"get_shipping_details_by_id/internal/controller"
	"github.com/gin-gonic/gin"
	"get_shipping_details_by_id/config/cors"
)

func SetupRouter(serviceId string) *gin.Engine {
	router := gin.Default()
	router.Use(cors.GetCORSConfig())
	//[ginshot-routes]
	router.GET("/get/:id", controller.NewGetShippingDetailsByIdController(service.NewGetShippingDetailsByIdServiceImpl(dbcontext.GetDBClient(), serviceId)).GetShippingDetailsById)
	//[HttpGET] Ping to get_shipping_details_by_id API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
	