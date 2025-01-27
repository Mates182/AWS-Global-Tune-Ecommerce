// auto-generated with ginshot
package service

import (
	"context"
	requests "create-billing-details-service/data/requests"
	responses "create-billing-details-service/data/responses"
	"fmt"
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
	mongoCollection := service.DBClient.Database("products").Collection("products")

	result, err := mongoCollection.InsertOne(context.Background(), request.Data.BillingDetails)
	if err != nil {
		return http.StatusInternalServerError, responses.Response{Message: "Error inserting billing details"}
	}

	fmt.Println(result.InsertedID)

	response := responses.Response{Message: "Billing details added successfully", Data: responses.ResponseData{BillingDetails: request.Data.BillingDetails}}
	return http.StatusOK, response
}
