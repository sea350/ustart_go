package auth

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/sea350/ustart_go/backend/auth/authpb"
)

// ChangePassword does what it does
func (rds *Redis) ChangePassword(ctx context.Context, req *authpb.ChangePasswordRequest) (*authpb.ChangePasswordResponse, error) {
	// authenticate
	_, err := rds.Authenticate(ctx, &authpb.AuthenticateRequest{
		UUID:      req.UUID,
		Challenge: req.Challenge,
	})
	if err != nil {
		return &authpb.ChangePasswordResponse{}, err
	}

	// encrypt the password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 255)
	if err != nil {
		return &authpb.ChangePasswordResponse{}, err
	}

	// set the new password
	_, err = rds.client.Set(uuidPwdPrefix+req.UUID.UUID, hashedPass, time.Duration(0)).Result()
	if err != nil {
		return &authpb.ChangePasswordResponse{}, err
	}

	return &authpb.ChangePasswordResponse{}, nil
}
