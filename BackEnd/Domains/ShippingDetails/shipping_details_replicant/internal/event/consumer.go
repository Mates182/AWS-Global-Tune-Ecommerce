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

// GetKafkaConsumer retorna una única instancia del consumidor (singleton)
func GetKafkaConsumer(brokers []string) (sarama.Consumer, error) {
	var err error
	consumerOnce.Do(func() {
		config := sarama.NewConfig()
		config.Consumer.Return.Errors = true
		consumerInstance, err = sarama.NewConsumer(brokers, config)
		if err != nil {
			log.Fatalf("Error creando el consumidor: %v", err)
		}
	})

	return consumerInstance, err
}

// CheckKafkaMessages lee los mensajes del topic y evita procesar su propio mensaje
func CheckKafkaMessages(kafkaBrokers []string, kafkaTopic string, serviceID string) bool {
	consumer, err := GetKafkaConsumer(kafkaBrokers)
	if err != nil {
		log.Println("Error obteniendo el consumer:", err)
		return false
	}

	partitionConsumer, err := consumer.ConsumePartition(kafkaTopic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Println("Error creando el partition consumer:", err)
		return false
	}
	defer partitionConsumer.Close()

	for msg := range partitionConsumer.Messages() {
		log.Printf("Mensaje recibido: %s\n", string(msg.Value))

		// Verificar si el mensaje fue enviado por esta misma instancia
		if string(msg.Key) == serviceID {
			log.Println("Mensaje propio detectado, ignorando...")
			continue
		}

		// Procesar mensaje
		return true
	}

	return false
}
