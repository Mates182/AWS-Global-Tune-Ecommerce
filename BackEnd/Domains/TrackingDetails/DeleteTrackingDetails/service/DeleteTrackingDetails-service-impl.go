// auto-generated with ginshot
package service

import (
	requests "delete-tracking-details-service/data/requests"
	responses "delete-tracking-details-service/data/responses"
	"net/http"
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
	response := responses.DeleteTrackingDetailsResponse{}
	return http.StatusOK, response
}
