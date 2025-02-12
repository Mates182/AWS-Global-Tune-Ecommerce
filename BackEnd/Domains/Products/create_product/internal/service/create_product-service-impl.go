package service

import (
	"encoding/json"
	"create_product/internal/event"
	"create_product/internal/secrets"
	message "create_product/internal/data/messages"
	requests "create_product/internal/data/requests"
	responses "create_product/internal/data/responses"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"fmt"
	"github.com/google/uuid"
	"context"
)
type CreateProductServiceImpl struct {
	// Add Components
	DBClient *mongo.Client
	ServiceId string
}

func NewCreateProductServiceImpl(dbClient *mongo.Client,serviceId string) CreateProductService {
	return &CreateProductServiceImpl{
		// Add Components
		DBClient: dbClient,
		ServiceId: serviceId,
	}
}

func (service *CreateProductServiceImpl) CreateProductHandler(request requests.CreateProductRequest) (int, responses.CreateProductResponse) {
	mongoCollection := service.DBClient.Database("products").Collection("products")
	if request.Product.ID != "" {
		idTag := "_id"
		existing := mongoCollection.FindOne(context.Background(), bson.M{idTag: request.Product.ID})
		if existing.Err() == nil {
			return http.StatusConflict, responses.CreateProductResponse{Message: "Product with the same ID already exists"}
		}
	}else {
		request.Product.ID = uuid.New().String()
	}

	result, err := mongoCollection.InsertOne(context.Background(), request.Product)
	if err != nil {
		return http.StatusInternalServerError, responses.CreateProductResponse{Message: "Error inserting Product"}
	}

	reply := &message.ProductTopicMessage{
		Service: service.ServiceId,
		Target:  "all",
		Product: request.Product,
		Action:  "create",
	}

	initMessage, _ := json.Marshal(reply)

	_ = event.PushToKafka("products", initMessage, secrets.GetKafkaBrokers())

	fmt.Println(result.InsertedID)

	response := responses.CreateProductResponse{Message: "Product created successfully", Product: request.Product}
	return http.StatusOK, response
}
