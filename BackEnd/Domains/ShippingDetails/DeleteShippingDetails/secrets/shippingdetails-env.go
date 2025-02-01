
package secrets

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetshippingdetailsDBURI() string {
	err := godotenv.Load()
	if err != nil {
		// TODO: remove on production
		fmt.Println("Error loading .env file, ignore if is on docker")
	}
	dbUri := os.Getenv("shippingdetails_URI")
	return dbUri
}
