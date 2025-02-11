package service

import (
	"encoding/json"
	"delete_shipping_details/internal/event"
	"delete_shipping_details/internal/secrets"
	message "delete_shipping_details/internal/data/messages"
	requests "delete_shipping_details/internal/data/requests"
	responses "delete_shipping_details/internal/data/responses"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"delete_shipping_details/internal/data/models"
	"context"
)
type DeleteShippingDetailsServiceImpl struct {
	// Add Components
	DBClient *mongo.Client
	ServiceId string
}

func NewDeleteShippingDetailsServiceImpl(dbClient *mongo.Client,serviceId string) DeleteShippingDetailsService {
	return &DeleteShippingDetailsServiceImpl{
		// Add Components
		DBClient: dbClient,
		ServiceId: serviceId,
	}
}

func (service *DeleteShippingDetailsServiceImpl) DeleteShippingDetailsHandler(request requests.DeleteShippingDetailsRequest) (int, responses.DeleteShippingDetailsResponse) {
	mongoCollection := service.DBClient.Database("shipping_details").Collection("shipping_details")

	idTag := "_id"
	filter := bson.M{idTag: request.ID}
	result, err := mongoCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return http.StatusInternalServerError, responses.DeleteShippingDetailsResponse{Message: "Error deleting ShippingDetails"}
	}
	if result.DeletedCount == 0 {
		return http.StatusNotFound, responses.DeleteShippingDetailsResponse{Message: "ShippingDetails not found"}
	}

	reply := &message.ShippingDetailsTopicMessage{
		Service: service.ServiceId,
		Target:  "all",
		ShippingDetails: models.ShippingDetails{ID: request.ID},
		Action:  "delete",
	}

	initMessage, _ := json.Marshal(reply)

	_ = event.PushToKafka("shipping_details", initMessage, secrets.GetKafkaBrokers())

	response := responses.DeleteShippingDetailsResponse{Message: "ShippingDetails deleted successfully", ID: request.ID}
	return http.StatusOK, response
}
