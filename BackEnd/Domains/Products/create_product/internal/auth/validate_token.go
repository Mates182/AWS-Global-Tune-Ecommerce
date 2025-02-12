package auth

import (
	"context"
	"time"
)

func ValidateTokenMiddleware(requiredRole string, token string) bool {
	req := &ValidateTokenRequest{
		Token:        token,
		RequiredRole: requiredRole,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := GetAuthClient().ValidateToken(ctx, req)
	if err != nil {
		return false
	}

	if !res.Valid {
		return false
	}
	return true

}
