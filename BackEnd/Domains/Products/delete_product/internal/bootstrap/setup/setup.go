package setup

import (
	"delete_product/internal/event"
	"delete_product/internal/secrets"
	"delete_product/internal/data/messages"
	"encoding/json"
)

// SetupService initializes the service and synchronizes data
func SetupService(serviceID string) error {
	brokers := secrets.GetKafkaBrokers()
	response := message.ReplicationTopicMessage{
		Service: serviceID,
		Action: "Instance Initialized",
	}

	initMessage, err := json.Marshal(response)
	if err != nil {
		return err
	}

	// Send the initialization message to Kafka
	err = event.PushToKafka("replication", initMessage, brokers)
	if err != nil {
		return err
	}

	return nil
}