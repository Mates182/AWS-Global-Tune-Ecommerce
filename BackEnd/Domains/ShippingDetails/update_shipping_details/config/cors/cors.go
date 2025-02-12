package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetCORSConfig() gin.HandlerFunc {
	corsConfig := cors.New(cors.Config{
		// Set to true to allow all origins (remove if you want to allow specific origins only)
		AllowAllOrigins: true, 

		// Uncomment and modify the line below to allow specific origins instead of all
		// AllowOrigins: []string{"http://localhost:80", "https://example.com"}, 

		// Define allowed HTTP methods (adjust according to your API needs)
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},

		// Specify the allowed headers (remove or add headers as required by your application)
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},

		// Set to true to allow credentials such as cookies or authorization headers
		AllowCredentials: true,
	})

	return corsConfig
}

