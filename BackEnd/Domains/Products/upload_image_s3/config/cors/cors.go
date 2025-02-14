package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetCORSConfig() gin.HandlerFunc {
	corsConfig := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * 3600, // 12 hours cache time
	})

	return corsConfig
}
