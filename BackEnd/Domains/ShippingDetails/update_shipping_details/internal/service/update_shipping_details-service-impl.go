package service

import (
	"encoding/json"
	"update_shipping_details/internal/event"
	"update_shipping_details/internal/secrets"
	message "update_shipping_details/internal/data/messages"
	requests "update_shipping_details/internal/data/requests"
	responses "update_shipping_details/internal/data/responses"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	
	"context"
)
type UpdateShippingDetailsServiceImpl struct {
	// Add Components
	DBClient *mongo.Client
	ServiceId string
}

func NewUpdateShippingDetailsServiceImpl(dbClient *mongo.Client,serviceId string) UpdateShippingDetailsService {
	return &UpdateShippingDetailsServiceImpl{
		// Add Components
		DBClient: dbClient,
		ServiceId: serviceId,
	}
}

func (service *UpdateShippingDetailsServiceImpl) UpdateShippingDetailsHandler(request requests.UpdateShippingDetailsRequest) (int, responses.UpdateShippingDetailsResponse) {
	mongoCollection := service.DBClient.Database("shipping_details").Collection("shipping_details")
	idTag := "_id"
	filter := bson.M{idTag: request.ShippingDetails.ID}
	update := bson.M{"$set": request.ShippingDetails}
	result, err := mongoCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return http.StatusInternalServerError, responses.UpdateShippingDetailsResponse{Message: "Error updating ShippingDetails"}
	}
	if result.MatchedCount == 0 {
		return http.StatusNotFound, responses.UpdateShippingDetailsResponse{Message: "ShippingDetails not found"}
	}
		
	reply := &message.ShippingDetailsTopicMessage{
		Service: service.ServiceId,
		Target:  "all",
		ShippingDetails: request.ShippingDetails,
		Action:  "update",
	}

	initMessage, _ := json.Marshal(reply)

	_ = event.PushToKafka("shipping_details", initMessage, secrets.GetKafkaBrokers())

	response := responses.UpdateShippingDetailsResponse{Message: "ShippingDetails updated successfully", ShippingDetails: request.ShippingDetails}
	return http.StatusOK, response
}
