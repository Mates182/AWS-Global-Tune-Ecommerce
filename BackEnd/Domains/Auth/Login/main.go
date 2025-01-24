package main

import (
	"login-service/controllers"
	_ "login-service/docs"
	"login-service/router"
	"login-service/secrets"
	"login-service/service"
)

// @title Login Service API
// @version 1.0
// @description RESTful API for user authentication in the GlobalTune E-Commerce platform.
// @description This API provides secure login functionality using JWT tokens and the Gin framework.

// @host 		localhost:80
// @BasePath 	/
func main() {
	jwtKey := secrets.GetJWTKey()
	loginService := service.NewLoginServiceImpl(jwtKey)
	loginController := controllers.NewLoginController(loginService)
	router := router.SetupRouter(loginController)
	router.Run("0.0.0.0:80")
}
