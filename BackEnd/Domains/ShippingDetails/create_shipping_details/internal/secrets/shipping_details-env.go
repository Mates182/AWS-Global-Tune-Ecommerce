
package secrets

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Getshipping_detailsDBURI() string {
	err := godotenv.Load()
	if err != nil {
		// TODO: remove on production
		fmt.Println("Error loading .env file, ignore if is on docker")
	}
	dbUri := os.Getenv("shipping_details_URI")
	return dbUri
}

func GetKafkaBrokers() []string {
	err := godotenv.Load()
	if err != nil {
		// TODO: remove on production
		fmt.Println("Error loading .env file, ignore if is on docker")
	}
	broker := []string{os.Getenv("shipping_details_BROKER")}
	return broker
}
