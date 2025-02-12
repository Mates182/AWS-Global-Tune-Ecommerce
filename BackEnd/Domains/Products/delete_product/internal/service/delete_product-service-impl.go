package service

import (
	"encoding/json"
	"delete_product/internal/event"
	"delete_product/internal/secrets"
	message "delete_product/internal/data/messages"
	requests "delete_product/internal/data/requests"
	responses "delete_product/internal/data/responses"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"delete_product/internal/data/models"
	"context"
)
type DeleteProductServiceImpl struct {
	// Add Components
	DBClient *mongo.Client
	ServiceId string
}

func NewDeleteProductServiceImpl(dbClient *mongo.Client,serviceId string) DeleteProductService {
	return &DeleteProductServiceImpl{
		// Add Components
		DBClient: dbClient,
		ServiceId: serviceId,
	}
}

func (service *DeleteProductServiceImpl) DeleteProductHandler(request requests.DeleteProductRequest) (int, responses.DeleteProductResponse) {
	mongoCollection := service.DBClient.Database("products").Collection("products")

	idTag := "_id"
	filter := bson.M{idTag: request.ID}
	result, err := mongoCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return http.StatusInternalServerError, responses.DeleteProductResponse{Message: "Error deleting Product"}
	}
	if result.DeletedCount == 0 {
		return http.StatusNotFound, responses.DeleteProductResponse{Message: "Product not found"}
	}

	reply := &message.ProductTopicMessage{
		Service: service.ServiceId,
		Target:  "all",
		Product: models.Product{ID: request.ID},
		Action:  "delete",
	}

	initMessage, _ := json.Marshal(reply)

	_ = event.PushToKafka("products", initMessage, secrets.GetKafkaBrokers())

	response := responses.DeleteProductResponse{Message: "Product deleted successfully", ID: request.ID}
	return http.StatusOK, response
}
