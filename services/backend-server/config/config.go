package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	StorageAPIKey string `yaml:"storageApiKey"`
	DatabaseURL   string `yaml:"databaseUrl"`
	IVEncryptionKey string `yaml:"ivEncryptionKey"`
}

func LoadConfig() (Config, error) {
	f, err := os.Open("config.yaml")
	if err != nil {
		return Config{}, err
	}
	defer f.Close()

	config := Config{}
	if err := yaml.NewDecoder(f).Decode(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
