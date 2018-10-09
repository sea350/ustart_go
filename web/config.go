package web

// Config is the configuration object for the web server
type Config struct {
	Port       string
	AssetsRoot string
}

// NewConfig returns a default Config object
func NewConfig() (*Config, error) {
	return &Config{}, nil
}
