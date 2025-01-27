package router

import (
	"create-tracking-details-service/dbcontext/trackingdetails"
	"create-tracking-details-service/service"
	"create-tracking-details-service/controller"
	"create-tracking-details-service/config/cors"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.GetCORSConfig())
	//[ginshot-routes]
	router.POST("/create", controller.NewCreateTrackingDetailsController(service.NewCreateTrackingDetailsServiceImpl(dbcontext.GetDBClient())).CreateTrackingDetails)
	// [HttpGET] Ping to create-tracking-details-service API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
