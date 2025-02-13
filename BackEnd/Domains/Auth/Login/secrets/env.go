package secrets

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, ignore if running in Docker")
	}
}

func GetJWTKey() []byte {
	LoadEnv()
	return []byte(os.Getenv("SECRET_KEY"))
}

func GetDBCredentials() (string, string, string, string, string) {
	LoadEnv()
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	return dbHost, dbUser, dbPassword, dbPort, dbName
}
