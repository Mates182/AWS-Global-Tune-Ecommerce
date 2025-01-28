package router

import (
	"get-tracking-details-by-id-service/dbcontext/trackingdetails"
	"get-tracking-details-by-id-service/service"
	"get-tracking-details-by-id-service/controller"
	"get-tracking-details-by-id-service/config/cors"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.GetCORSConfig())
	//[ginshot-routes]
	router.GET("/get/:id", controller.NewGetTrackingDetailsByIdController(service.NewGetTrackingDetailsByIdServiceImpl(dbcontext.GetDBClient())).GetTrackingDetailsById)
	// [HttpGET] Ping to get-tracking-details-by-id-service API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
