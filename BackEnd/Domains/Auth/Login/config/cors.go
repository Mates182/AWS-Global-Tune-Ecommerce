package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetCORSConfig() gin.HandlerFunc {
	corsConfig := cors.New(cors.Config{
		AllowOrigins:     []string{"http://52.90.4.29:80"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	})
	return corsConfig
}
