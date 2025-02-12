package main

import (
	"fmt"
	"upload_image_s3/router"
)

func main() {
	r := router.SetupRouter()

	// Start the server
	fmt.Println("Server started on port :80")
	r.Run("0.0.0.0:80")
}
