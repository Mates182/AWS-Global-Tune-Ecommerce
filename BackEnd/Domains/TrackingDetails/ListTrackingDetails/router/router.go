package router

import (
	"list-tracking-details/dbcontext/trackingdetails"
	"list-tracking-details/service"
	"list-tracking-details/controller"
	"github.com/gin-gonic/gin"
	"list-tracking-details/config/cors"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.GetCORSConfig())
	//[ginshot-routes]
	router.GET("/list", controller.NewListTrackingDetailsController(service.NewListTrackingDetailsServiceImpl(dbcontext.GetDBClient())).ListTrackingDetails)
	//[HttpGET] Ping to list-tracking-details API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
	