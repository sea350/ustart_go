package auth

import (
	"net"

	"github.com/go-redis/redis"
	"github.com/sea350/ustart_go/backend/auth/authpb"
	"google.golang.org/grpc"
)

const (
	uuidPwdPrefix   = "_uuid_pwd:"
	uuidUnamePrefix = "_uuid_uname:"
	unameUUIDPrefix = "_uuid_pwd:"
)

// Redis is an implementation of the auth service defined in service.proto
// We maintain the following indexes, with this service:
// uuid : hashedPwd (auth)
// uuid : username (identity)
// username : uuid (identity-lookup)
type Redis struct {
	client     *redis.Client
	gRPCServer *grpc.Server
	gRPCPort   string
}

// New returns a new Redis auth server, which is started with Start() and Stopped with Shutdown()
func New(cfg *Config) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})

	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}

	rds := &Redis{
		client:   client,
		gRPCPort: cfg.GRPCPort,
	}

	srv := grpc.NewServer()
	authpb.RegisterAuthServer(srv, rds)
	rds.gRPCServer = srv

	return rds, nil
}

// Start begins the grpc-accessible, redis-backed auth service
func (rds *Redis) Start() error {
	lis, err := net.Listen("tcp", ":"+rds.gRPCPort)
	if err != nil {
		panic(err)
	}

	return rds.gRPCServer.Serve(lis)
}

// Shutdown gracefully stops the redis-backed auth service
func (rds *Redis) Shutdown() error {
	rds.gRPCServer.GracefulStop()
	return nil
}
