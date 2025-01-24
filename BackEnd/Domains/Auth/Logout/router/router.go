package router

import (
	"logout-service/config"
	"logout-service/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(logoutController *controllers.LogoutController) *gin.Engine {
	router := gin.Default()

	router.Use(config.GetCORSConfig())

	router.POST("logout", logoutController.PostLogout)
	return router
}
