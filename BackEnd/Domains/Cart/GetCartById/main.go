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

	client = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0,
	})
	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("Failed to connect to Redis: %s", err.Error())
		return
	}
	fmt.Println(ping)

	router := gin.Default()
	// RESTful routes
	router.GET("cart/:id", getCartByID)

	router.Run("0.0.0.0:80")

}

// GET: Get cart by ID
func getCartByID(c *gin.Context) {
	id := c.Param("id")
	cart := "cart:" + id
	if client == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Client not initialized"})
		return
	}
	val, err := client.HGetAll(context.Background(), cart).Result()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Cart not found"})
		return
	}
	if len(val) == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Cart is empty"})
		return
	}
	c.IndentedJSON(http.StatusOK, val)

}
