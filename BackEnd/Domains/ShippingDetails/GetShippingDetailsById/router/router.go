package router

import (
	"github.com/gin-gonic/gin"
	"get-shipping-details-by-id/config/cors"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.GetCORSConfig())
	//[ginshot-routes]
	//[HttpGET] Ping to get-shipping-details-by-id API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
	