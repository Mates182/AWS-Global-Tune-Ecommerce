package router

import (
	"logout-service/config"
	"logout-service/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(logoutController *controllers.LogoutController) *gin.Engine {
	router := gin.Default()

	router.Use(config.GetCORSConfig())
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("logout", logoutController.PostLogout)
	return router
}
