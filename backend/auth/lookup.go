package auth

import (
	"context"

	"github.com/sea350/ustart_go/backend/auth/authpb"
)

// Lookup does what it does
func (rds *Redis) Lookup(ctx context.Context, req *authpb.LookupRequest) (*authpb.LookupResponse, error) {
	uuid, err := rds.client.Get(unameUUIDPrefix + req.Username).Result()
	if err != nil {
		return nil, err
	}

	return &authpb.LookupResponse{
		UUID: &authpb.UUID{UUID: uuid},
	}, nil
}
