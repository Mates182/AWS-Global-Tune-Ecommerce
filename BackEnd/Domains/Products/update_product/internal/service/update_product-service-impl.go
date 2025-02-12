package service

import (
	"encoding/json"
	"update_product/internal/event"
	"update_product/internal/secrets"
	message "update_product/internal/data/messages"
	requests "update_product/internal/data/requests"
	responses "update_product/internal/data/responses"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	
	"context"
)
type UpdateProductServiceImpl struct {
	// Add Components
	DBClient *mongo.Client
	ServiceId string
}

func NewUpdateProductServiceImpl(dbClient *mongo.Client,serviceId string) UpdateProductService {
	return &UpdateProductServiceImpl{
		// Add Components
		DBClient: dbClient,
		ServiceId: serviceId,
	}
}

func (service *UpdateProductServiceImpl) UpdateProductHandler(request requests.UpdateProductRequest) (int, responses.UpdateProductResponse) {
	mongoCollection := service.DBClient.Database("products").Collection("products")
	idTag := "_id"
	filter := bson.M{idTag: request.Product.ID}
	update := bson.M{"$set": request.Product}
	result, err := mongoCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return http.StatusInternalServerError, responses.UpdateProductResponse{Message: "Error updating Product"}
	}
	if result.MatchedCount == 0 {
		return http.StatusNotFound, responses.UpdateProductResponse{Message: "Product not found"}
	}
		
	reply := &message.ProductTopicMessage{
		Service: service.ServiceId,
		Target:  "all",
		Product: request.Product,
		Action:  "update",
	}

	initMessage, _ := json.Marshal(reply)

	_ = event.PushToKafka("products", initMessage, secrets.GetKafkaBrokers())

	response := responses.UpdateProductResponse{Message: "Product updated successfully", Product: request.Product}
	return http.StatusOK, response
}
