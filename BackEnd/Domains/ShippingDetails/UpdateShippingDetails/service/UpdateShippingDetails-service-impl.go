package service

import (
	requests "update-shipping-details/data/requests"
	responses "update-shipping-details/data/responses"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	
	"context"
)
type UpdateShippingDetailsServiceImpl struct {
	// Add Components
	DBClient *mongo.Client
}

func NewUpdateShippingDetailsServiceImpl(dbClient *mongo.Client) UpdateShippingDetailsService {
	return &UpdateShippingDetailsServiceImpl{
		// Add Components
		DBClient: dbClient,
	}
}

func (service *UpdateShippingDetailsServiceImpl) UpdateShippingDetailsHandler(request requests.UpdateShippingDetailsRequest) (int, responses.UpdateShippingDetailsResponse) {
	mongoCollection := service.DBClient.Database("shippingdetails").Collection("shippingdetails")
	filter := bson.M{"ID": request.ShippingDetails.ID}
	update := bson.M{"$set": request.ShippingDetails}
	result, err := mongoCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return http.StatusInternalServerError, responses.UpdateShippingDetailsResponse{Message: "Error updating ShippingDetails"}
	}
	if result.MatchedCount == 0 {
		return http.StatusNotFound, responses.UpdateShippingDetailsResponse{Message: "ShippingDetails not found"}
	}

	response := responses.UpdateShippingDetailsResponse{Message: "ShippingDetails updated successfully", ShippingDetails: request.ShippingDetails}
	return http.StatusOK, response
}
