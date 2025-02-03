package secrets

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetKafkaBrokers() []string {
	err := godotenv.Load()
	if err != nil {
		// TODO: remove on production
		fmt.Println("Error loading .env file, ignore if is on docker")
	}
	broker := []string{os.Getenv("shipping_details_BROKER")}
	return broker
}
