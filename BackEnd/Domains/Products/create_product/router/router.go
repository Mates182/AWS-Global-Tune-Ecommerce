package router

import (
	"create_product/internal/repository"
	"create_product/internal/service"
	"create_product/internal/controller"
	"github.com/gin-gonic/gin"
	"create_product/config/cors"
)

func SetupRouter(serviceId string) *gin.Engine {
	router := gin.Default()
	router.Use(cors.GetCORSConfig())
	//[ginshot-routes]
	router.POST("/create", controller.NewCreateProductController(service.NewCreateProductServiceImpl(dbcontext.GetDBClient(), serviceId)).CreateProduct)
	//[HttpGET] Ping to create_product API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
	