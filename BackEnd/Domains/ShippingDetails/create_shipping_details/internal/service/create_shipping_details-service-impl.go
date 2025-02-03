package service

import (
	"encoding/json"
	"create_shipping_details/internal/event"
	"create_shipping_details/internal/secrets"
	message "create_shipping_details/internal/data/messages"
	requests "create_shipping_details/internal/data/requests"
	responses "create_shipping_details/internal/data/responses"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"fmt"
	"github.com/google/uuid"
	"context"
)
type CreateShippingDetailsServiceImpl struct {
	// Add Components
	DBClient *mongo.Client
	ServiceId string
}

func NewCreateShippingDetailsServiceImpl(dbClient *mongo.Client,serviceId string) CreateShippingDetailsService {
	return &CreateShippingDetailsServiceImpl{
		// Add Components
		DBClient: dbClient,
		ServiceId: serviceId,
	}
}

func (service *CreateShippingDetailsServiceImpl) CreateShippingDetailsHandler(request requests.CreateShippingDetailsRequest) (int, responses.CreateShippingDetailsResponse) {
	mongoCollection := service.DBClient.Database("shipping_details").Collection("shipping_details")
	if request.ShippingDetails.ID != "" {
		idTag := "_id"
		existing := mongoCollection.FindOne(context.Background(), bson.M{idTag: request.ShippingDetails.ID})
		if existing.Err() == nil {
			return http.StatusConflict, responses.CreateShippingDetailsResponse{Message: "ShippingDetails with the same ID already exists"}
		}
	}else {
		request.ShippingDetails.ID = uuid.New().String()
	}

	result, err := mongoCollection.InsertOne(context.Background(), request.ShippingDetails)
	if err != nil {
		return http.StatusInternalServerError, responses.CreateShippingDetailsResponse{Message: "Error inserting ShippingDetails"}
	}

	reply := &message.ShippingDetailsTopicMessage{
		Service: service.ServiceId,
		Target:  "all",
		ShippingDetails: request.ShippingDetails,
		Action:  "create",
	}

	initMessage, _ := json.Marshal(reply)

	_ = event.PushToKafka("shipping_details", initMessage, secrets.GetKafkaBrokers())

	fmt.Println(result.InsertedID)

	response := responses.CreateShippingDetailsResponse{Message: "ShippingDetails created successfully", ShippingDetails: request.ShippingDetails}
	return http.StatusOK, response
}
