// auto-generated with ginshot
package service

import (
	"context"
	requests "get-tracking-details-by-id-service/data/requests"
	responses "get-tracking-details-by-id-service/data/responses"
	"get-tracking-details-by-id-service/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GetTrackingDetailsByIdServiceImpl struct {
	// Add Components
	DBClient *mongo.Client
}

func NewGetTrackingDetailsByIdServiceImpl(dbClient *mongo.Client) GetTrackingDetailsByIdService {
	return &GetTrackingDetailsByIdServiceImpl{
		// Add Components
		DBClient: dbClient,
	}
}

func (service *GetTrackingDetailsByIdServiceImpl) GetTrackingDetailsByIdHandler(request requests.GetTrackingDetailsByIdRequest) (int, responses.GetTrackingDetailsByIdResponse) {
	mongoCollection := service.DBClient.Database("trackingdetails").Collection("trackingdetails")

	var trackingDetail models.TrackingDetails
	err := mongoCollection.FindOne(context.Background(), bson.M{"OrderID": request.OrderID}).Decode(&trackingDetail)
	if err == mongo.ErrNoDocuments {
		return http.StatusNotFound, responses.GetTrackingDetailsByIdResponse{Message: "Tracking Details not found"}
	}
	if err != nil {
		return http.StatusInternalServerError, responses.GetTrackingDetailsByIdResponse{Message: "Error fetching tracking details"}
	}

	response := responses.GetTrackingDetailsByIdResponse{Message: "Tracking Details retrieved successfully", TrackingDetails: trackingDetail}
	return http.StatusOK, response
}
