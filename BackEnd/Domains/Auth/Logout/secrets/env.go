package secrets

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetJWTKey() []byte {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, ignore if is on docker")
	}
	jwtKey := []byte(os.Getenv("SECRET_KEY"))
	return jwtKey
}
