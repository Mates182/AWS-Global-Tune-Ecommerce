package service

import (
	requests "get_shipping_details_by_id/internal/data/requests"
	responses "get_shipping_details_by_id/internal/data/responses"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"get_shipping_details_by_id/internal/data/models"
	"context"
)
type GetShippingDetailsByIdServiceImpl struct {
	// Add Components
	DBClient *mongo.Client
	ServiceId string
}

func NewGetShippingDetailsByIdServiceImpl(dbClient *mongo.Client,serviceId string) GetShippingDetailsByIdService {
	return &GetShippingDetailsByIdServiceImpl{
		// Add Components
		DBClient: dbClient,
		ServiceId: serviceId,
	}
}

func (service *GetShippingDetailsByIdServiceImpl) GetShippingDetailsByIdHandler(request requests.GetShippingDetailsByIdRequest) (int, responses.GetShippingDetailsByIdResponse) {
	mongoCollection := service.DBClient.Database("shipping_details").Collection("shipping_details")

	var ShippingDetails models.ShippingDetails
	idTag := "_id"
	err := mongoCollection.FindOne(context.Background(), bson.M{idTag: request.ID}).Decode(&ShippingDetails)
	if err == mongo.ErrNoDocuments {
		return http.StatusNotFound, responses.GetShippingDetailsByIdResponse{Message: "ShippingDetails not found"}
	}
	if err != nil {
		return http.StatusInternalServerError, responses.GetShippingDetailsByIdResponse{Message: "Error fetching ShippingDetails"}
	}

	response := responses.GetShippingDetailsByIdResponse{Message: "ShippingDetails retrieved successfully", ShippingDetails: ShippingDetails}
	return http.StatusOK, response
}
