package secrets

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetDBURI() string {
	err := godotenv.Load()
	if err != nil {
		// TODO: remove on production
		fmt.Println("Error loading .env file, ignore if is on docker")
	}
	dbUri := os.Getenv("BILLING_DETAILS_DB_URI")
	return dbUri
}

func GetCORSOrigins() string {
	err := godotenv.Load()
	if err != nil {
		// TODO: remove on production
		fmt.Println("Error loading .env file, ignore if is on docker")
	}
	corsOrigins := os.Getenv("CORS_ORIGINS")
	return corsOrigins
}
