package auth

import (
	"context"
	"testing"

	"github.com/sea350/ustart_go/backend/auth/authpb"
)

func TestRedis_Register_Simple(t *testing.T) {
	// create a redis client instance
	rds, err := New(NewConfig())
	if err != nil {
		t.Fatalf("Failed to create the auth server {%+v}", err)
	}

	regResp, err := rds.Register(context.Background(), &authpb.RegisterRequest{
		Username: "test_username",
		Password: "test_password",
	})

	if err != nil {
		t.Fatalf("Failed to register a user {%+v}", err)
	}

	// test all indices exist

	// uname_uuid
	lookupResp, err := rds.Lookup(context.Background(), &authpb.LookupRequest{
		Username: "test_username",
	})
	if err != nil {
		t.Fatalf("Failed to lookup user {%v}", err)
	}
	if lookupResp.UUID.UUID != regResp.UUID.UUID {
		t.Fatalf("uuids {%v} vs {%v} do not match", lookupResp.UUID.UUID, regResp.UUID.UUID)
	}

	// uuid_password
	_, err = rds.Authenticate(context.Background(), &authpb.AuthenticateRequest{
		UUID:      regResp.UUID,
		Challenge: "test_password",
	})
	if err != nil {
		t.Fatalf("Failed to authenticate user {%v}", err)
	}

	// uuid_username
	// TODO(adam) check this by doing chuname

}

func TestRedis_Register_FailsOnUnameExisting(t *testing.T) {
	// create a redis client instance
	rds, err := New(NewConfig())
	if err != nil {
		t.Fatalf("Failed to create the auth server {%+v}", err)
	}

	_, err = rds.Register(context.Background(), &authpb.RegisterRequest{
		Username: "test_username2",
		Password: "test_password",
	})
	if err != nil {
		t.Fatalf("Failed to register a user {%+v}", err)
	}

	_, err = rds.Register(context.Background(), &authpb.RegisterRequest{
		Username: "test_username2",
		Password: "test_password",
	})
	if err != errDuplicateUname {
		t.Fatalf("Expected errDuplicateUname, instead got {%v}", err)
	}
}
