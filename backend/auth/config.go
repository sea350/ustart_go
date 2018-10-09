package auth

// Config determines the runtime behavior of the redis-backed auth server
type Config struct {
	RedisAddr string
	GRPCPort  string
}

// NewConfig returns a default config object
func NewConfig() *Config {
	return &Config{
		RedisAddr: "localhost:6379",
		GRPCPort:  "1234",
	}
}
