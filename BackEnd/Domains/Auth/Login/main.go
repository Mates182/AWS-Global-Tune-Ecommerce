package main

import (
	"login-service/router"
)

func main() {
	router := router.NewRouter()
	router.Run("0.0.0.0:80")
}
