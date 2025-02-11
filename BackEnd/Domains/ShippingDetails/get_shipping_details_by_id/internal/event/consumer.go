package event

import (
	"log"
	"sync"

	"github.com/IBM/sarama"
)

var (
	consumerInstance sarama.Consumer
	consumerOnce     sync.Once
)

func GetKafkaConsumer(brokers []string) (sarama.Consumer, error) {
	var err error
	consumerOnce.Do(func() {
		config := sarama.NewConfig()
		config.Consumer.Return.Errors = true
		consumerInstance, err = sarama.NewConsumer(brokers, config)
		if err != nil {
			log.Fatalf("Error creating consumer: %v", err)
		}
	})

	return consumerInstance, err
}
