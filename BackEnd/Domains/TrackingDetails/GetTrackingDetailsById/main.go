package main

import (
	"fmt"
	"get-tracking-details-by-id-service/router"
)

func main() {
	fmt.Println("get-tracking-details-by-id-service API started!")
	router := router.SetupRouter()
	router.Run("0.0.0.0:80")
}
