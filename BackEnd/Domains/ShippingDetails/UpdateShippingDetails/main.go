package main

import (
	"update-shipping-details/router"
	"fmt"
)

func main() {
	fmt.Println("update-shipping-details API started!")
	router := router.SetupRouter()
	router.Run("0.0.0.0:80")
}
		