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
	dbUri := os.Getenv("ORDERS_DB_URI")
	return dbUri
}
func GetDBPassword() string {
	err := godotenv.Load()
	if err != nil {
		// TODO: remove on production
		fmt.Println("Error loading .env file, ignore if is on docker")
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	return dbPassword
}
func GetAuthServiceURI() string {
	err := godotenv.Load()
	if err != nil {
		// TODO: remove on production
		fmt.Println("Error loading .env file, ignore if is on docker")
	}
	authServiceURI := os.Getenv("AUTH_SERVICE_URI")
	return authServiceURI
}
