package router

import (
	"login-service/controllers"

	"login-service/config"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.Use(config.GetCORSConfig())

	router.POST("login", controllers.PostLogin)
	return router
}
