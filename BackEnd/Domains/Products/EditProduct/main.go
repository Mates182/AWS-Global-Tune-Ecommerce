package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"products-rest-api/auth"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://0.0.0.0:3000"},
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router.PATCH("edit/:id", patchProductByID)

	router.Run("0.0.0.0:80")

}

func patchProductByID(c *gin.Context) {
	endpoint := os.Getenv("AUTH_SERVICE_URL")
	rpcConn, err := grpc.NewClient(endpoint, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer rpcConn.Close()

	authClient := auth.NewAuthServiceClient(rpcConn)
	token, err := c.Cookie("token")
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Missing or invalid token"})
		return
	}
	fmt.Println(token)
	req := &auth.ValidateTokenRequest{
		Token:        token,
		RequiredRole: "admin",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := authClient.ValidateToken(ctx, req)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error validating token"})
		fmt.Println(err)
		return
	}

	if !res.Valid {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token or role"})
		return
	}

	id := c.Param("id")
	var updates map[string]interface{}

	if err := c.BindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if mongoCollection == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not initialized"})
		return
	}

	filter := bson.M{"sku": id}

	update := bson.M{"$set": updates}

	result, err := mongoCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}
