package main

import (
	"delete_shipping_details/internal/bootstrap/setup"
	"delete_shipping_details/internal/event"
	"delete_shipping_details/router"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func main() {
	// Generate a unique service ID (UUID)
	serviceID := uuid.New().String()
	log.Println("Starting service with UUID: " + serviceID + "\n")

	// Initialize Kafka consumer in a separate function
	go event.StartKafkaConsumer(serviceID)

	// Initialize service and synchronize data
	setup.SetupService(serviceID)

	// Start the API server
	fmt.Println("DeleteShippingDetails API started!")
	router := router.SetupRouter(serviceID)
	router.Run("0.0.0.0:80")
}