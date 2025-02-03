
package dbcontext

import (
	"context"
	"create_shipping_details/internal/secrets"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var clientInstance *mongo.Client
var clientOnce sync.Once

func GetDBClient() *mongo.Client {
	clientOnce.Do(func() {
		endpoint := secrets.Getshipping_detailsDBURI()
		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(endpoint))
		if err != nil {
			panic(err)
		}
		fmt.Println("Connected to shipping_details Database Server")

		err = client.Ping(context.Background(), readpref.Primary())
		if err != nil {
			panic(err)
		}
		fmt.Println("Pong")
		clientInstance = client
	})

	return clientInstance
}
		