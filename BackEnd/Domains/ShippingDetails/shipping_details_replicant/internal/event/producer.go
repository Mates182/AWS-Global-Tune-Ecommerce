package event

import (
	"log"
	"sync"

	"github.com/IBM/sarama"
)

var (
	producerInstance sarama.SyncProducer
	producerOnce     sync.Once
)

// GetKafkaProducer retorna una única instancia del productor (singleton)
func GetKafkaProducer(brokers []string) (sarama.SyncProducer, error) {
	var err error
	producerOnce.Do(func() {
		config := sarama.NewConfig()
		config.Producer.Return.Successes = true
		config.Producer.RequiredAcks = sarama.WaitForAll
		config.Producer.Retry.Max = 5

		producerInstance, err = sarama.NewSyncProducer(brokers, config)
		if err != nil {
			log.Fatalf("Error creando el producer: %v", err)
		}
	})

	return producerInstance, err
}

// PushToKafka envía un mensaje al topic de Kafka usando el singleton
func PushToKafka(topic string, message []byte, brokers []string) error {
	producer, err := GetKafkaProducer(brokers)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	_, _, err = producer.SendMessage(msg)
	return err
}
