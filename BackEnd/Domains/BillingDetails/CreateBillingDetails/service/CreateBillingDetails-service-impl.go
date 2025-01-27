// auto-generated with ginshot
package service

import (
	requests "create-billing-details-service/data/requests"
	responses "create-billing-details-service/data/responses"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type CreateBillingDetailsServiceImpl struct {
	// Add Components
	DBClient *mongo.Client
}

func NewCreateBillingDetailsServiceImpl(dbClient *mongo.Client) CreateBillingDetailsService {
	return &CreateBillingDetailsServiceImpl{
		// Add Components
		DBClient: dbClient,
	}
}

func (service *CreateBillingDetailsServiceImpl) CreateBillingDetailsHandler(request requests.Request) (int, responses.Response) {
	response := responses.Response{}
	return http.StatusOK, response
}
