package main

import (
	"login-service/controllers"
	"login-service/router"
	"login-service/secrets"
	"login-service/service"
)

func main() {
	jwtKey := secrets.GetJWTKey()
	loginService := service.NewLoginServiceImpl(jwtKey)
	loginController := controllers.NewLoginController(loginService)
	router := router.SetupRouter(loginController)
	router.Run("0.0.0.0:80")
}
