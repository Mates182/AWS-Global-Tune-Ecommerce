package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Product struct {
	ID          string   `json:"id" bson:"id"`
	Title       string   `json:"title" bson:"title"`
	Description string   `json:"description" bson:"description"`
	Category    string   `json:"category" bson:"category"`
	Price       float64  `json:"price" bson:"price"`
	Stock       int      `json:"stock" bson:"stock"`
	Tags        []string `json:"tags" bson:"tags"`
	Brand       string   `json:"brand" bson:"brand"`
	Sku         string   `json:"sku" bson:"sku"`
	Weight      float64  `json:"weight" bson:"weight"`
	Warranty    string   `json:"warranty" bson:"warranty"`
	Thumbnail   string   `json:"thumbnail" bson:"thumbnail"`
	Images      []string `json:"images" bson:"images"`
}

func newMongoClient(endpoint string) *mongo.Client {
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(endpoint))
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")
	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	fmt.Println("Pong")
	return mongoClient
}

var mongoCollection *mongo.Collection

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, ignore if is on docker")
	}
	endpoint := os.Getenv("MONGO_URI")
	client := newMongoClient(endpoint)
	defer client.Disconnect(context.Background())
	mongoCollection = client.Database("globaltune_products").Collection("products")

	router := gin.Default()
	router.POST("/products/create", createProduct)

	router.Run("0.0.0.0:80")

}

func createProduct(c *gin.Context) {
	if mongoCollection == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Client not initialized"})
		return
	}

	// TODO: authenticate admin rights

	var product bson.M

	if err := c.BindJSON(&product); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	if product["sku"] != nil {
		existing := mongoCollection.FindOne(context.Background(), bson.M{"sku": product["sku"]})
		if existing.Err() == nil {
			c.IndentedJSON(http.StatusConflict, gin.H{"message": "Product with the same SKU already exists"})
			return
		}
	}

	result, err := mongoCollection.InsertOne(context.Background(), product)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error inserting product"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"insertedID": result.InsertedID})
}
