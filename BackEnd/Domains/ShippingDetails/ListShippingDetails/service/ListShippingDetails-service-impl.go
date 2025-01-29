package service

import (
	"context"
	requests "list-shipping-details/data/requests"
	responses "list-shipping-details/data/responses"
	"list-shipping-details/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ListShippingDetailsServiceImpl struct {
	// Add Components
	DBClient *mongo.Client
}

func NewListShippingDetailsServiceImpl(dbClient *mongo.Client) ListShippingDetailsService {
	return &ListShippingDetailsServiceImpl{
		// Add Components
		DBClient: dbClient,
	}
}

func (service *ListShippingDetailsServiceImpl) ListShippingDetailsHandler(request requests.ListShippingDetailsRequest) (int, responses.ListShippingDetailsResponse) {
	mongoCollection := service.DBClient.Database("shippingdetails").Collection("shippingdetails")

	cursor, err := mongoCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return http.StatusInternalServerError, responses.ListShippingDetailsResponse{Message: "Error fetching ShippingDetails"}
	}
	defer cursor.Close(context.Background())

	var ShippingDetails []models.ShippingDetails
	if err := cursor.All(context.Background(), &ShippingDetails); err != nil {
		return http.StatusInternalServerError, responses.ListShippingDetailsResponse{Message: "Error decoding ShippingDetails"}
	}

	response := responses.ListShippingDetailsResponse{Message: "All ShippingDetails retrieved successfully", ShippingDetails: ShippingDetails}
	return http.StatusOK, response
}
