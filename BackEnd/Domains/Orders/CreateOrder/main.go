package main

import (
	"context"
	"create-order-service/dbcontext"
	"fmt"
)

func main() {
	fmt.Println("GlobalTune / Create Order Service")
	client := dbcontext.GetDBClient()
	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("Failed to connect to DB: %s\n", err.Error())
		return
	}
	fmt.Printf("Ping: %s\n", ping)
}
