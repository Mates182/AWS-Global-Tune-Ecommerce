package service

import (
	"context"
	requests "get-shipping-details-by-id/data/requests"
	responses "get-shipping-details-by-id/data/responses"
	"get-shipping-details-by-id/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GetShippingDetailsByIdServiceImpl struct {
	// Add Components
	DBClient *mongo.Client
}

func NewGetShippingDetailsByIdServiceImpl(dbClient *mongo.Client) GetShippingDetailsByIdService {
	return &GetShippingDetailsByIdServiceImpl{
		// Add Components
		DBClient: dbClient,
	}
}

func (service *GetShippingDetailsByIdServiceImpl) GetShippingDetailsByIdHandler(request requests.GetShippingDetailsByIdRequest) (int, responses.GetShippingDetailsByIdResponse) {
	mongoCollection := service.DBClient.Database("shippingdetails").Collection("shippingdetails")

	var ShippingDetails models.ShippingDetails
	err := mongoCollection.FindOne(context.Background(), bson.M{"ID": request.ShippingDetails.ID}).Decode(&ShippingDetails)
	if err == mongo.ErrNoDocuments {
		return http.StatusNotFound, responses.GetShippingDetailsByIdResponse{Message: "ShippingDetails not found"}
	}
	if err != nil {
		return http.StatusInternalServerError, responses.GetShippingDetailsByIdResponse{Message: "Error fetching ShippingDetails"}
	}

	response := responses.GetShippingDetailsByIdResponse{Message: "ShippingDetails retrieved successfully", ShippingDetails: ShippingDetails}
	return http.StatusOK, response
}
