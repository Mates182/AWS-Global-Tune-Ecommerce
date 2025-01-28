// auto-generated with ginshot
package service

import (
	"context"
	"net/http"
	requests "update-tracking-details-service/data/requests"
	responses "update-tracking-details-service/data/responses"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UpdateTrackingDetailsServiceImpl struct {
	// Add Components
	DBClient *mongo.Client
}

func NewUpdateTrackingDetailsServiceImpl(dbClient *mongo.Client) UpdateTrackingDetailsService {
	return &UpdateTrackingDetailsServiceImpl{
		// Add Components
		DBClient: dbClient,
	}
}

func (service *UpdateTrackingDetailsServiceImpl) UpdateTrackingDetailsHandler(request requests.UpdateTrackingDetailsRequest) (int, responses.UpdateTrackingDetailsResponse) {
	mongoCollection := service.DBClient.Database("trackingdetails").Collection("trackingdetails")

	filter := bson.M{"OrderID": request.TrackingDetails.OrderID}
	update := bson.M{"$set": request.TrackingDetails}

	result, err := mongoCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return http.StatusInternalServerError, responses.UpdateTrackingDetailsResponse{Message: "Error updating tracking details"}
	}
	if result.MatchedCount == 0 {
		return http.StatusNotFound, responses.UpdateTrackingDetailsResponse{Message: "Tracking Details not found"}
	}

	response := responses.UpdateTrackingDetailsResponse{Message: "Tracking Details updated successfully", TrackingDetails: request.TrackingDetails}
	return http.StatusOK, response
}
