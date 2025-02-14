package main

import (
	"login-service/controllers"
	"login-service/database"
	_ "login-service/docs"
	"login-service/router"
	"login-service/secrets"
	"login-service/service"
)

// @title Login Service API
// @version 1.0
// @description RESTful API for user authentication in the GlobalTune E-Commerce platform.
// @description This API provides secure login functionality using JWT tokens and the Gin framework.

// @host 		3.84.60.208
// @BasePath 	/
func main() {
	// Load secrets
	secrets.LoadEnv()
	// Get singleton DB instance
	db := database.GetDBInstance()
	defer db.Close() // Ensure DB is closed when the app exits
	jwtKey := secrets.GetJWTKey()
	loginService := service.NewLoginServiceImpl(jwtKey, db)
	loginController := controllers.NewLoginController(loginService)
	router := router.SetupRouter(loginController)
	router.RunTLS("0.0.0.0:443", "/root/cert.pem", "/root/key.pem")
}
