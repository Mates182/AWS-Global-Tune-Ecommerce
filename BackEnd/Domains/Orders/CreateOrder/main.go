package main

import (
	"create-order-service/controllers"
	"create-order-service/dbcontext"
	"create-order-service/router"
	"create-order-service/service"
	"fmt"
)

func main() {
	fmt.Println("GlobalTune / Create Order Service")
	client := dbcontext.GetDBClient()
	createOrderService := service.NewCreateOrderServiceImpl(client)
	createOrderController := controllers.NewCreateOrderController(createOrderService)
	router := router.SetupRouter(createOrderController)
	router.Run("0.0.0.0:80")

}
