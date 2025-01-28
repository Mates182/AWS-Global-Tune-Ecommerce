package router

import (
	"update-tracking-details-service/dbcontext/trackingdetails"
	"update-tracking-details-service/service"
	"update-tracking-details-service/controller"
	"update-tracking-details-service/config/cors"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.GetCORSConfig())
	//[ginshot-routes]
	router.PATCH("/update", controller.NewUpdateTrackingDetailsController(service.NewUpdateTrackingDetailsServiceImpl(dbcontext.GetDBClient())).UpdateTrackingDetails)
	// [HttpGET] Ping to update-tracking-details-service API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
