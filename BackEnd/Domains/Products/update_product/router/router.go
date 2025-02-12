package router

import (
	"update_product/internal/repository"
	"update_product/internal/service"
	"update_product/internal/controller"
	"github.com/gin-gonic/gin"
	"update_product/config/cors"
)

func SetupRouter(serviceId string) *gin.Engine {
	router := gin.Default()
	router.Use(cors.GetCORSConfig())
	//[ginshot-routes]
	router.PATCH("/update", controller.NewUpdateProductController(service.NewUpdateProductServiceImpl(dbcontext.GetDBClient(), serviceId)).UpdateProduct)
	//[HttpGET] Ping to update_product API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
	