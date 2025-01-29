package main

import (
	"delete-shipping-details/router"
	"fmt"
)

func main() {
	fmt.Println("delete-shipping-details API started!")
	router := router.SetupRouter()
	router.Run("0.0.0.0:80")
}
		