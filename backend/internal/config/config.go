package config

type Config struct {
}

func LoadConfig() (*Config, error) {

	cfg := Config{}

	return &cfg, nil
}
