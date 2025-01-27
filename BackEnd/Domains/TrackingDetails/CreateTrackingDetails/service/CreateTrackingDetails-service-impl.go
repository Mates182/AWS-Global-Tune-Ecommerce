// auto-generated with ginshot
package service

import (
	requests "create-tracking-details-service/data/requests"
	responses "create-tracking-details-service/data/responses"
	"net/http"
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
	response := responses.CreateTrackingDetailsResponse{}
	return http.StatusOK, response
}
