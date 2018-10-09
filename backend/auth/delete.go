package auth

import (
	"context"
	"fmt"

	"github.com/sea350/ustart_go/backend/auth/authpb"
)

// Delete does what it does
func (rds *Redis) Delete(ctx context.Context, req *authpb.DeleteRequest) (*authpb.DeleteResponse, error) {
	// authenticate
	_, err := rds.Authenticate(ctx, &authpb.AuthenticateRequest{
		UUID:      req.UUID,
		Challenge: req.Challenge,
	})
	if err != nil {
		return &authpb.DeleteResponse{}, err
	}

	// get the uname, so we can clear its index
	uname, err := rds.client.Get(uuidUnamePrefix + req.UUID.UUID).Result()
	if err != nil {
		return &authpb.DeleteResponse{}, err
	}

	// delete indices
	numDeleted, err := rds.client.Del(uuidPwdPrefix+req.UUID.UUID, uuidUnamePrefix+req.UUID.UUID, unameUUIDPrefix+uname).Result()
	if err != nil {
		return &authpb.DeleteResponse{}, err
	}
	if numDeleted != 3 {
		return &authpb.DeleteResponse{}, fmt.Errorf("Deleted %d indices instead of 3", numDeleted)
	}

	return &authpb.DeleteResponse{}, nil
}
