package main

import (
	"create-billing-details-service/router"
	"fmt"
)

func main() {
	fmt.Println("GlobalTune / Create Billing Details Service")
	router := router.SetupRouter()
	router.Run("0.0.0.0:80")
}
