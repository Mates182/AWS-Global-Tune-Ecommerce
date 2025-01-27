package router

import (
	"delete-tracking-details-service/config/cors"
	"delete-tracking-details-service/controller"
	dbcontext "delete-tracking-details-service/dbcontext/trackingdetails"
	"delete-tracking-details-service/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.GetCORSConfig())
	//[ginshot-routes]
	router.DELETE("/delete/:id", controller.NewDeleteTrackingDetailsController(service.NewDeleteTrackingDetailsServiceImpl(dbcontext.GetDBClient())).DeleteTrackingDetails)
	// [HttpGET] Ping to delete-tracking-details-service API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
