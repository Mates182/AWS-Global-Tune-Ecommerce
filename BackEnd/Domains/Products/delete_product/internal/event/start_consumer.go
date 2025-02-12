package event

import (
	"delete_product/config/kafka"
	message "delete_product/internal/data/messages"
	dbcontext "delete_product/internal/repository"
	"delete_product/internal/secrets"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

// StartKafkaConsumer initializes and runs the Kafka consumer in a separate goroutine
func StartKafkaConsumer(serviceID string) {
	topic := "products"
	msgCnt := 0

	// Initialize Kafka consumer
	consumer, err := GetKafkaConsumer(secrets.GetKafkaBrokers())
	if err != nil {
		log.Fatalf("Error obtaining Kafka consumer: %v", err)
	}
	worker, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	defer worker.Close()

	isServiceActive := false
	fmt.Println("Kafka consumer started")

	replicator := dbcontext.NewReplicator(*dbcontext.NewProductRepository(dbcontext.GetDBClient().Database("products")))

	// Handle OS signals for graceful shutdown
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-worker.Errors():
			fmt.Println("Kafka Error:", err)
		case msg := <-worker.Messages():
			msgCnt++
			fmt.Printf("Received message %d: | Topic(%s)\n", msgCnt, msg.Topic)

			var ProductMsg message.ProductTopicMessage
			err := json.Unmarshal(msg.Value, &ProductMsg)
			if err != nil {
				fmt.Println("Error parsing Kafka message:", err)
				continue
			}

		
			// Log parsed message
			fmt.Printf("Parsed Kafka Message: %+v\n", ProductMsg)
			if ProductMsg.Target == serviceID {
				fmt.Printf("Message targeted to this service: %+v\n", ProductMsg)
				// Handle database updates
				if ProductMsg.Action == "db-update" {
					fmt.Println("Database update detected, syncing offset...")
					kafka.GetOffsetManager().SetOffset(msg.Offset)
					isServiceActive = true
					fmt.Println("Service is now active")
				}
			} else if ProductMsg.Target == "all" && ProductMsg.Service != serviceID {
				if isServiceActive {
					fmt.Printf("Processing update for Product ID: %s\n", ProductMsg.Product.ID)
					replicator.Replicate(ProductMsg.Product, ProductMsg.Action)
				}
			}

		case <-sigchan:
			fmt.Println("Interrupt detected, shutting down consumer...")
			return
		}
	}
}
