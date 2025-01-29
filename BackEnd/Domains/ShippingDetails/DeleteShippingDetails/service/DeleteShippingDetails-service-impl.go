package service

import (
	requests "delete-shipping-details/data/requests"
	responses "delete-shipping-details/data/responses"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	
	"context"
)
type DeleteShippingDetailsServiceImpl struct {
	// Add Components
	DBClient *mongo.Client
}

func NewDeleteShippingDetailsServiceImpl(dbClient *mongo.Client) DeleteShippingDetailsService {
	return &DeleteShippingDetailsServiceImpl{
		// Add Components
		DBClient: dbClient,
	}
}

func (service *DeleteShippingDetailsServiceImpl) DeleteShippingDetailsHandler(request requests.DeleteShippingDetailsRequest) (int, responses.DeleteShippingDetailsResponse) {
	mongoCollection := service.DBClient.Database("shippingdetails").Collection("shippingdetails")

	filter := bson.M{"ID": request.ShippingDetails.ID}
	result, err := mongoCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return http.StatusInternalServerError, responses.DeleteShippingDetailsResponse{Message: "Error deleting ShippingDetails"}
	}
	if result.DeletedCount == 0 {
		return http.StatusNotFound, responses.DeleteShippingDetailsResponse{Message: "ShippingDetails not found"}
	}

	response := responses.DeleteShippingDetailsResponse{Message: "ShippingDetails deleted successfully", ShippingDetails: request.ShippingDetails}
	return http.StatusOK, response
}
