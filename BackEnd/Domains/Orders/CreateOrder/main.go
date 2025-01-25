package main

import (
	"create-order-service/controllers"
	"create-order-service/dbcontext"
	_ "create-order-service/docs"
	"create-order-service/router"
	"create-order-service/service"
	"fmt"
)

// @title Create Order Service API
// @version 1.0
// @description RESTful API to handle the creation of customer orders in the GlobalTune E-Commerce platform.
// @description This API create orders securely using JWT tokens and the Gin framework.

// @host localhost:80
// @BasePath /
func main() {
	fmt.Println("GlobalTune / Create Order Service")
	client := dbcontext.GetDBClient()
	createOrderService := service.NewCreateOrderServiceImpl(client)
	createOrderController := controllers.NewCreateOrderController(createOrderService)
	router := router.SetupRouter(createOrderController)
	router.Run("0.0.0.0:80")

}
