package auth

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/sea350/ustart_go/backend/auth/authpb"
)

// Register does what it does
func (rds *Redis) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	// make sure username is not already taken
	_, err := rds.Lookup(ctx, &authpb.LookupRequest{
		Username: req.Username,
	})
	if err != redis.Nil {
		return &authpb.RegisterResponse{}, errDuplicateUname
	}

	// generate uuids until we get one that is not being used
	var uuidStr string
	for {
		uuid, err := uuid.NewUUID()
		if err != nil {
			return &authpb.RegisterResponse{}, err
		}

		uuidStr = uuid.String()
		_, err = rds.client.Get(uuidUnamePrefix + uuidStr).Result()
		if err == redis.Nil {
			break
		}
	}

	// encrypt the password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return &authpb.RegisterResponse{}, err
	}

	// create all necessary indices
	good, err := rds.client.MSetNX(
		uuidUnamePrefix+uuidStr, req.Username,
		uuidPwdPrefix+uuidStr, string(hashedPass),
		unameUUIDPrefix+req.Username, uuidStr,
	).Result()
	if err != nil {
		return &authpb.RegisterResponse{}, err
	}
	if !good {
		return &authpb.RegisterResponse{}, fmt.Errorf("Failed to create indices, because some of them already existed")
	}

	return &authpb.RegisterResponse{
		UUID: &authpb.UUID{UUID: uuidStr},
	}, nil
}
