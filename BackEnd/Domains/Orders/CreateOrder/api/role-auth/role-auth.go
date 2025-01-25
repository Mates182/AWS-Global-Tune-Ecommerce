package roleauth

import (
	"context"
	"create-order-service/api/role-auth/auth"
	"create-order-service/secrets"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

func AuthenticateUser(token string) bool {
	endpoint := secrets.GetAuthServiceURI()
	fmt.Println(endpoint)
	rpcConn, err := grpc.NewClient(endpoint, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("Failed to connect to server: %v", err)
	}
	defer rpcConn.Close()

	authClient := auth.NewAuthServiceClient(rpcConn)

	req := &auth.ValidateTokenRequest{
		Token:        token,
		RequiredRole: "admin",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := authClient.ValidateToken(ctx, req)
	if err != nil {
		fmt.Printf("Error validating token: %v", err)
		return false
	}

	if !res.Valid {
		fmt.Printf("Invalid token or role: %v", err)
		return false
	}
	return true
}
