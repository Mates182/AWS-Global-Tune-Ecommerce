
package secrets

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GettrackingdetailsDBURI() string {
	err := godotenv.Load()
	if err != nil {
		// TODO: remove on production
		fmt.Println("Error loading .env file, ignore if is on docker")
	}
	dbUri := os.Getenv("trackingdetails_URI")
	return dbUri
}
