package config

func GenConfig(elastic_url string, username string, password string) Config {
	return Config{ELASTIC_URL: elastic_url, USERNAME: username, PASSWORD: password}
}


