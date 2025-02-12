package auth

import (
	"log"
	"os"
	"sync"

	"google.golang.org/grpc"
)

var (
	authClient AuthServiceClient
	once       sync.Once
)

func GetAuthClient() AuthServiceClient {
	once.Do(func() {
		endpoint := os.Getenv("AUTH_SERVICE_URL")
		conn, err := grpc.NewClient(endpoint, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Failed to connect to authentication service: %v", err)
		}

		authClient = NewAuthServiceClient(conn)
	})
	return authClient
}
