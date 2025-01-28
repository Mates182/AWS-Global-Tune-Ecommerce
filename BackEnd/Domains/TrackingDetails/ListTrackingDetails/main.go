package main

import (
	"list-tracking-details/router"
	"fmt"
)

func main() {
	fmt.Println("list-tracking-details API started!")
	router := router.SetupRouter()
	router.Run("0.0.0.0:80")
}
		