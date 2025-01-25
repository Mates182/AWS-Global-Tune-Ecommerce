package router

import (
	"create-billing-details-service/config"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(config.GetCORSConfig())
	return router
}
