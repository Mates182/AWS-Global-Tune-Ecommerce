package service

import (
	"context"
	requests "list-tracking-details/data/requests"
	responses "list-tracking-details/data/responses"
	"list-tracking-details/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ListTrackingDetailsServiceImpl struct {
	// Add Components
	DBClient *mongo.Client
}

func NewListTrackingDetailsServiceImpl(dbClient *mongo.Client) ListTrackingDetailsService {
	return &ListTrackingDetailsServiceImpl{
		// Add Components
		DBClient: dbClient,
	}
}

func (service *ListTrackingDetailsServiceImpl) ListTrackingDetailsHandler(request requests.ListTrackingDetailsRequest) (int, responses.ListTrackingDetailsResponse) {
	mongoCollection := service.DBClient.Database("trackingdetails").Collection("trackingdetails")

	cursor, err := mongoCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return http.StatusInternalServerError, responses.ListTrackingDetailsResponse{Message: "Error fetching TrackingDetails"}
	}
	defer cursor.Close(context.Background())

	var TrackingDetails []models.TrackingDetails
	if err := cursor.All(context.Background(), &TrackingDetails); err != nil {
		return http.StatusInternalServerError, responses.ListTrackingDetailsResponse{Message: "Error decoding TrackingDetails"}
	}

	response := responses.ListTrackingDetailsResponse{Message: "All TrackingDetails retrieved successfully", TrackingDetails: TrackingDetails}
	return http.StatusOK, response
}
