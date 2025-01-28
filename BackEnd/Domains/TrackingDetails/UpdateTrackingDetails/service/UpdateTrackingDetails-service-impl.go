// auto-generated with ginshot
package service

import (
	requests "update-tracking-details-service/data/requests"
	responses "update-tracking-details-service/data/responses"
	"net/http"
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
	response := responses.UpdateTrackingDetailsResponse{}
	return http.StatusOK, response
}
