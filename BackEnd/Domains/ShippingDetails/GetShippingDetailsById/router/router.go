package router

import (
	"get-shipping-details-by-id/dbcontext/shippingdetails"
	"get-shipping-details-by-id/service"
	"get-shipping-details-by-id/controller"
	"github.com/gin-gonic/gin"
	"get-shipping-details-by-id/config/cors"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.GetCORSConfig())
	//[ginshot-routes]
	router.GET("/get/:id", controller.NewGetShippingDetailsByIdController(service.NewGetShippingDetailsByIdServiceImpl(dbcontext.GetDBClient())).GetShippingDetailsById)
	//[HttpGET] Ping to get-shipping-details-by-id API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
	