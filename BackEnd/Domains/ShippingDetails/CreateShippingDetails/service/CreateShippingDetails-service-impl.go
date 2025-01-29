package service

import (
	"context"
	requests "create-shipping-details/data/requests"
	responses "create-shipping-details/data/responses"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateShippingDetailsServiceImpl struct {
	// Add Components
	DBClient *mongo.Client
}

func NewCreateShippingDetailsServiceImpl(dbClient *mongo.Client) CreateShippingDetailsService {
	return &CreateShippingDetailsServiceImpl{
		// Add Components
		DBClient: dbClient,
	}
}

func (service *CreateShippingDetailsServiceImpl) CreateShippingDetailsHandler(request requests.CreateShippingDetailsRequest) (int, responses.CreateShippingDetailsResponse) {
	mongoCollection := service.DBClient.Database("shippingdetails").Collection("shippingdetails")
	if request.ShippingDetails.ID != "" {
		existing := mongoCollection.FindOne(context.Background(), bson.M{"ID": request.ShippingDetails.ID})
		if existing.Err() == nil {
			return http.StatusConflict, responses.CreateShippingDetailsResponse{Message: "ShippingDetails with the same ID already exists"}
		}
	}

	result, err := mongoCollection.InsertOne(context.Background(), request.ShippingDetails)
	if err != nil {
		return http.StatusInternalServerError, responses.CreateShippingDetailsResponse{Message: "Error inserting ShippingDetails"}
	}

	fmt.Println(result.InsertedID)

	response := responses.CreateShippingDetailsResponse{Message: "ShippingDetails created successfully", ShippingDetails: request.ShippingDetails}
	return http.StatusOK, response
}
