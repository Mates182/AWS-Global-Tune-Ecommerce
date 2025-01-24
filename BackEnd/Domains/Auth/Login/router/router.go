package router

import (
	"login-service/controllers"

	"login-service/config"

	"github.com/gin-gonic/gin"
)

func SetupRouter(loginController *controllers.LoginController) *gin.Engine {
	router := gin.Default()

	router.Use(config.GetCORSConfig())

	router.POST("login", loginController.PostLogin)
	return router
}
