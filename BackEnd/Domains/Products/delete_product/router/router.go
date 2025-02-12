package router

import (
	"delete_product/internal/repository"
	"delete_product/internal/service"
	"delete_product/internal/controller"
	"github.com/gin-gonic/gin"
	"delete_product/config/cors"
)

func SetupRouter(serviceId string) *gin.Engine {
	router := gin.Default()
	router.Use(cors.GetCORSConfig())
	//[ginshot-routes]
	router.DELETE("/delete/:id", controller.NewDeleteProductController(service.NewDeleteProductServiceImpl(dbcontext.GetDBClient(), serviceId)).DeleteProduct)
	//[HttpGET] Ping to delete_product API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
	