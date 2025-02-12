
package secrets

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetproductsDBURI() string {
	err := godotenv.Load()
	if err != nil {
		// TODO: remove on production
		fmt.Println("Error loading .env file, ignore if is on docker")
	}
	dbUri := os.Getenv("products_URI")
	return dbUri
}

func GetKafkaBrokers() []string {
	err := godotenv.Load()
	if err != nil {
		// TODO: remove on production
		fmt.Println("Error loading .env file, ignore if is on docker")
	}
	broker := []string{os.Getenv("products_BROKER")}
	return broker
}
