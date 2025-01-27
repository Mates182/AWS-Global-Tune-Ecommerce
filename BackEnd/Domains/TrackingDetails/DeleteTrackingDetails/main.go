package main

import (
	"delete-tracking-details-service/router"
	"fmt"
)

func main() {
	fmt.Println("delete-tracking-details-service API started!")
	router := router.SetupRouter()
	router.Run("0.0.0.0:80")
}
