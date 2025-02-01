package main

import (
	"fmt"
	"update-tracking-details-service/router"
)

func main() {
	fmt.Println("update-tracking-details-service API started!")
	router := router.SetupRouter()
	router.Run("0.0.0.0:80")
}
