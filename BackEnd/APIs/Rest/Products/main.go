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
	router.GET("products/", getAllProducts)
	router.GET("products/:id", getProductByID)
	router.POST("products", createProduct)
	router.DELETE("products/:id", deleteProductByID)

	router.Run("0.0.0.0:8000")

}

func getAllProducts(c *gin.Context) {
	if mongoCollection == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Client not initialized"})
		return
	}
	cursor, err := mongoCollection.Find(context.Background(), bson.D{})
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Error retrieving products"})
		return
	}
	defer cursor.Close(context.Background())
	var products []bson.M
	if err := cursor.All(context.Background(), &products); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error reading products"})
		return
	}

	if len(products) == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "No products found"})
		return
	}

	c.IndentedJSON(http.StatusOK, products)

}

func getProductByID(c *gin.Context) {
	id := c.Param("id")

	if mongoCollection == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Client not initialized"})
		return
	}

	filter := bson.M{"sku": id}

	var product bson.M
	err := mongoCollection.FindOne(context.Background(), filter).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving product"})
		return
	}

	c.IndentedJSON(http.StatusOK, product)
}

func createProduct(c *gin.Context) {
	if mongoCollection == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Client not initialized"})
		return
	}

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

func deleteProductByID(c *gin.Context) {
	if mongoCollection == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Client not initialized"})
		return
	}

	sku := c.Param("id")
	if sku == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "SKU is required"})
		return
	}

	result, err := mongoCollection.DeleteOne(context.Background(), bson.M{"sku": sku})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error deleting product"})
		return
	}

	if result.DeletedCount == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// TODO: implement PUT/PATCH
