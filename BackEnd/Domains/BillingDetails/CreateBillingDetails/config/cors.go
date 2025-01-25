package config

import (
	"create-billing-details-service/secrets"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetCORSConfig() gin.HandlerFunc {
	corsConfig := cors.New(cors.Config{
		AllowOrigins:     []string{secrets.GetCORSOrigins()},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	return corsConfig
}
