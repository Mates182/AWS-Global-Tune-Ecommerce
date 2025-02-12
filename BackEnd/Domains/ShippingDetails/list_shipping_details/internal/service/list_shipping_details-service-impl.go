package service

import (
	
	responses "list_shipping_details/internal/data/responses"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"list_shipping_details/internal/data/models"
	"context"
)
type ListShippingDetailsServiceImpl struct {
	// Add Components
	DBClient *mongo.Client
	ServiceId string
}

func NewListShippingDetailsServiceImpl(dbClient *mongo.Client,serviceId string) ListShippingDetailsService {
	return &ListShippingDetailsServiceImpl{
		// Add Components
		DBClient: dbClient,
		ServiceId: serviceId,
	}
}

func (service *ListShippingDetailsServiceImpl) ListShippingDetailsHandler() (int, responses.ListShippingDetailsResponse) {
	mongoCollection := service.DBClient.Database("shipping_details").Collection("shipping_details")

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
