package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var client *redis.Client

// Cart struct
type Cart struct {
	ID    string `json:"id"`
	Items []struct {
		ProductID string `json:"product_id"`
		Quantity  int    `json:"quantity"`
	} `json:"items"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, ignore if is on docker")
	}
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisOptions := &redis.Options{
		Addr: redisAddr,
		DB:   0,
	}
	if redisPassword != "" {
		redisOptions.Password = redisPassword
	}

	client = redis.NewClient(redisOptions)
	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("Failed to connect to Redis: %s\n", err.Error())
		return
	}
	fmt.Println(ping)

	router := gin.Default()
	// RESTful routes
	router.POST("cart/", postCart)

	router.Run("0.0.0.0:80")

}

// POST: Add a new cart or Create if don't exists
func postCart(c *gin.Context) {
	var cart Cart
	if err := c.BindJSON(&cart); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}
	if client == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Redis client not initialized"})
		return
	}

	cartKey := "cart:" + cart.ID
	for _, item := range cart.Items {
		err := client.HSet(context.Background(), cartKey, item.ProductID, item.Quantity).Err()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to add item to cart"})
			return
		}
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Cart created successfully"})
}
