package main

import (
	"context"
	"create-billing-details-service/dbcontetxt"
	"create-billing-details-service/router"
	"fmt"
)

func main() {
	fmt.Println("GlobalTune / Create Billing Details Service")
	dbclient := dbcontetxt.GetDBClient()
	defer dbclient.Disconnect(context.Background())
	router := router.SetupRouter()
	router.Run("0.0.0.0:80")
}
