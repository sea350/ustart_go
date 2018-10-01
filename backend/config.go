package backend

// Config is used to determine the runtime behavior of backend.Server
type Config struct {
	Port string
}

// NewConfig returns a new config object with default params
func NewConfig() *Config {
	return &Config{}
}
