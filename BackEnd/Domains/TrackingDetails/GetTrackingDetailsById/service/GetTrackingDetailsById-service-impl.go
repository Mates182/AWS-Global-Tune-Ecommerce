// auto-generated with ginshot
package service

import (
	requests "get-tracking-details-by-id-service/data/requests"
	responses "get-tracking-details-by-id-service/data/responses"
	"net/http"
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
	response := responses.GetTrackingDetailsByIdResponse{}
	return http.StatusOK, response
}
