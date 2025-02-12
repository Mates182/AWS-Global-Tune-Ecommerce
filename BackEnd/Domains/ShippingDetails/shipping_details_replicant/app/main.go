package main

import (
	"encoding/json"
	"fmt"
	"models-replicant/config/cors"
	message "models-replicant/internal/data/messages"
	"models-replicant/internal/data/models"
	"models-replicant/internal/event"
	"models-replicant/internal/secrets"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
)

func main() {
	topic := "replication"
	msgCnt := 0

	// 1. Create a new consumer and start it.
	worker, err := event.GetKafkaConsumer(secrets.GetKafkaBrokers())
	if err != nil {
		panic(err)
	}

	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	fmt.Println("Consumer started")
	// 2. Handle OS signals - used to stop the process.
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// 3. Create a Goroutine to run the consumer / worker.
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println("Error:", err)
			case msg := <-consumer.Messages():
				msgCnt++
				ProcessMsg(msg, msgCnt)
				// TODO: implement db replications
			case <-sigchan:
				fmt.Println("Interrupt detected")
			}
		}
	}()

	router := gin.Default()
	router.Use(cors.GetCORSConfig())
	router.Run("0.0.0.0:80")

	// 4. Close the consumer on exit.
	if err := worker.Close(); err != nil {
		panic(err)
	}
}

func ProcessMsg(msg *sarama.ConsumerMessage, index int) {
	fmt.Printf("Received message Count %d: | Topic(%s)\n", index, string(msg.Topic))

	var serviceMsg message.ReplicationTopicMessage
	err := json.Unmarshal(msg.Value, &serviceMsg)
	if err != nil {
		fmt.Println("Error parsing message:", err)
		return
	}
	fmt.Printf("Service: %s | Message: %s\n", serviceMsg.Service, serviceMsg.Action)

	if index >= 1 {
		fmt.Println("Activating service for the first message...")

		responseModel := models.ShippingDetails{}
		response := message.ShippingDetailsTopicMessage{
			Service:         "models-replicant",
			Action:          "db-update",
			ShippingDetails: responseModel,
			Target:          serviceMsg.Service,
		}

		initMessage, err := json.Marshal(response)
		if err != nil {
			fmt.Println("Error serializing response:", err)
			return
		}

		brokers := secrets.GetKafkaBrokers()
		topic := "shipping_details"
		err = event.PushToKafka(topic, initMessage, brokers)
		if err != nil {
			fmt.Println("Error sending message to Kafka:", err)
			return
		}

		fmt.Println("Response message sent to Kafka!")
	}
}

func ConnectConsumer(brokers []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	return sarama.NewConsumer(brokers, config)
}
