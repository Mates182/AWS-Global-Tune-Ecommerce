package main

import (
	"create-shipping-details/router"
	"fmt"
)

func main() {
	fmt.Println("create-shipping-details API started!")
	router := router.SetupRouter()
	router.Run("0.0.0.0:80")
}
		