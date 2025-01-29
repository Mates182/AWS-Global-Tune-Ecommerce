package main

import (
	"list-shipping-details/router"
	"fmt"
)

func main() {
	fmt.Println("list-shipping-details API started!")
	router := router.SetupRouter()
	router.Run("0.0.0.0:80")
}
		