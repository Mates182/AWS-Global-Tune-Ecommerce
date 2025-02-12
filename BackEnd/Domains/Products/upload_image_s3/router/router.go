package router

import (
	"upload_image_s3/config/cors"
	"upload_image_s3/internal/awsutils"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.GetCORSConfig())
	router.POST("/upload", awsutils.UploadImageHandler)
	//[HttpGET] Ping to delete_product API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
