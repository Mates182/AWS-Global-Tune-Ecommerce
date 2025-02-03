package main

import (
	"get_shipping_details_by_id/internal/bootstrap/setup"
	"get_shipping_details_by_id/internal/event"
	"get_shipping_details_by_id/router"
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
	fmt.Println("GetShippingDetailsById API started!")
	router := router.SetupRouter(serviceID)
	router.Run("0.0.0.0:80")
}