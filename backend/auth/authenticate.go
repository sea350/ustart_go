package auth

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/sea350/ustart_go/backend/auth/authpb"
)

// Authenticate performs a lookup
func (rds *Redis) Authenticate(ctx context.Context, req *authpb.AuthenticateRequest) (*authpb.AuthenticateResponse, error) {
	hashedPwd, err := rds.client.Get(uuidPwdPrefix + req.UUID.UUID).Result()
	if err != nil {
		return &authpb.AuthenticateResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(req.Challenge))
	if err != nil {
		return &authpb.AuthenticateResponse{}, err
	}

	return &authpb.AuthenticateResponse{}, nil
}
