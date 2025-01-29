package main

import (
	"get-shipping-details-by-id/router"
	"fmt"
)

func main() {
	fmt.Println("get-shipping-details-by-id API started!")
	router := router.SetupRouter()
	router.Run("0.0.0.0:80")
}
		