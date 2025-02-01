// auto-generated with ginshot
package service

import (
	"context"
	requests "create-tracking-details-service/data/requests"
	responses "create-tracking-details-service/data/responses"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateTrackingDetailsServiceImpl struct {
	// Add Components
	DBClient *mongo.Client
}

func NewCreateTrackingDetailsServiceImpl(dbClient *mongo.Client) CreateTrackingDetailsService {
	return &CreateTrackingDetailsServiceImpl{
		// Add Components
		DBClient: dbClient,
	}
}

func (service *CreateTrackingDetailsServiceImpl) CreateTrackingDetailsHandler(request requests.CreateTrackingDetailsRequest) (int, responses.CreateTrackingDetailsResponse) {
	mongoCollection := service.DBClient.Database("trackingdetails").Collection("trackingdetails")
	if request.TrackingDetails.OrderID != "" {
		existing := mongoCollection.FindOne(context.Background(), bson.M{"OrderID": request.TrackingDetails.OrderID})
		if existing.Err() == nil {
			return http.StatusConflict, responses.CreateTrackingDetailsResponse{Message: "Tracking Details with the same OrderID already exists"}
		}
	}

	result, err := mongoCollection.InsertOne(context.Background(), request.TrackingDetails)
	if err != nil {
		return http.StatusInternalServerError, responses.CreateTrackingDetailsResponse{Message: "Error inserting tracking details"}
	}

	fmt.Println(result.InsertedID)

	response := responses.CreateTrackingDetailsResponse{Message: "Tracking Details created successfully", TrackingDetails: request.TrackingDetails}
	return http.StatusOK, response
}
