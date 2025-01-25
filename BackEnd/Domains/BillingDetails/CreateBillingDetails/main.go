package main

import (
	"context"
	"create-billing-details-service/dbcontetxt"
	"fmt"
)

func main() {
	fmt.Println("GlobalTune / Create Billing Details Service")
	dbclient := dbcontetxt.GetDBClient()
	defer dbclient.Disconnect(context.Background())
}
