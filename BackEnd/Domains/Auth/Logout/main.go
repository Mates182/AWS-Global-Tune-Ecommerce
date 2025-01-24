package main

import (
	"logout-service/controllers"
	"logout-service/router"
	"logout-service/secrets"
	"logout-service/service"
)

func main() {
	jwtKey := secrets.GetJWTKey()
	logoutService := service.NewLogoutServiceImpl(jwtKey)
	logoutController := controllers.NewLogoutController(logoutService)
	router := router.SetupRouter(logoutController)
	router.Run("0.0.0.0:80")
}
