// auto-generated with ginshot
package service

import (
	"context"
	requests "delete-tracking-details-service/data/requests"
	responses "delete-tracking-details-service/data/responses"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeleteTrackingDetailsServiceImpl struct {
	// Add Components
	DBClient *mongo.Client
}

func NewDeleteTrackingDetailsServiceImpl(dbClient *mongo.Client) DeleteTrackingDetailsService {
	return &DeleteTrackingDetailsServiceImpl{
		// Add Components
		DBClient: dbClient,
	}
}

func (service *DeleteTrackingDetailsServiceImpl) DeleteTrackingDetailsHandler(request requests.DeleteTrackingDetailsRequest) (int, responses.DeleteTrackingDetailsResponse) {
	mongoCollection := service.DBClient.Database("trackingdetails").Collection("trackingdetails")

	filter := bson.M{"OrderID": request.TrackingDetails.OrderID}
	result, err := mongoCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return http.StatusInternalServerError, responses.DeleteTrackingDetailsResponse{Message: "Error deleting tracking details"}
	}
	if result.DeletedCount == 0 {
		return http.StatusNotFound, responses.DeleteTrackingDetailsResponse{Message: "Tracking Details not found"}
	}

	response := responses.DeleteTrackingDetailsResponse{Message: "Tracking Details deleted successfully"}
	return http.StatusOK, response
}
