package dbcontext

import (
	"context"
	"create-order-service/secrets"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func GetDBClient() *redis.Client {

	dbURI := secrets.GetDBURI()
	dbPassword := secrets.GetDBPassword()
	dbOptions := &redis.Options{
		Addr: dbURI,
		DB:   0,
	}
	if dbPassword != "" {
		dbOptions.Password = dbPassword
	}

	client := redis.NewClient(dbOptions)
	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("Failed to connect to DB: %s\n", err.Error())
		return nil
	}
	fmt.Printf("Ping: %s\n", ping)
	return client
}
