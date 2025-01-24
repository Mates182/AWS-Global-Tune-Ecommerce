package main

import (
	"logout-service/controllers"
	_ "logout-service/docs"
	"logout-service/router"
	"logout-service/secrets"
	"logout-service/service"
)

// @Title Logout Service API
// @Version 1.0
// @Description RESTful API for user logout in the GlobalTune E-Commerce platform.
// @Description This API handles user session termination securely using JWT tokens and the Gin framework.

// @Host localhost:80
// @BasePath /
func main() {
	jwtKey := secrets.GetJWTKey()
	logoutService := service.NewLogoutServiceImpl(jwtKey)
	logoutController := controllers.NewLogoutController(logoutService)
	router := router.SetupRouter(logoutController)
	router.Run("0.0.0.0:80")
}
