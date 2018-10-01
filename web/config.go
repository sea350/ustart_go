package web

type Config struct {
	Port       string
	AssetsRoot string
}

func NewConfig() (*Config, error) {
	return &Config{}, nil
}
