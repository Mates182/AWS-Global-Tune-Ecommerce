package router

import (
	"delete-shipping-details/dbcontext/shippingdetails"
	"delete-shipping-details/service"
	"delete-shipping-details/controller"
	"github.com/gin-gonic/gin"
	"delete-shipping-details/config/cors"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.GetCORSConfig())
	//[ginshot-routes]
	router.DELETE("/delete/:id", controller.NewDeleteShippingDetailsController(service.NewDeleteShippingDetailsServiceImpl(dbcontext.GetDBClient())).DeleteShippingDetails)
	//[HttpGET] Ping to delete-shipping-details API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
	