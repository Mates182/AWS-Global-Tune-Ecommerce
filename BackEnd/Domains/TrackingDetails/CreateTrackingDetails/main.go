package main

import (
	"create-tracking-details-service/router"
	"fmt"
)

func main() {
	fmt.Println("create-tracking-details-service API started!")
	router := router.SetupRouter()
	router.Run("0.0.0.0:80")
}
