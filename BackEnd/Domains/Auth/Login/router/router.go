package router

import (
	"login-service/config"
	"login-service/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(loginController *controllers.LoginController) *gin.Engine {
	router := gin.Default()

	router.Use(config.GetCORSConfig())
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("login", loginController.PostLogin)
	return router
}
